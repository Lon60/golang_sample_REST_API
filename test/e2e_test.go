package e2e_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"golang_sample/internal/config"
	"golang_sample/internal/domain/demo"
	"golang_sample/internal/domain/user"
)

// setupTestServer initializes the app with an in-memory SQLite DB using shared cache and returns the HTTP test server and DB reference.
func setupTestServer(jwtSecret string) (*httptest.Server, *gorm.DB) {
	// Use SQLite in-memory mode with shared cache
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Auto migrate the models
	err = db.AutoMigrate(&demo.Demo{}, &user.User{})
	if err != nil {
		panic(err)
	}

	// Initialize demo and user parts
	demoRepo := demo.NewDemoRepository(db)
	demoService := demo.NewDemoService(demoRepo.Repository)
	demoHandler := demo.NewDemoHandler(demoService)

	userRepo := user.NewRepository(db)
	userHandler := user.NewUserHandler(userRepo, jwtSecret)

	// Setup Gin router
	r := gin.Default()

	api := r.Group("/api")
	{
		// Demo routes secured with JWT
		demo.RegisterRoutes(api, demoHandler, jwtSecret)

		// User routes (registration & login)
		user.RegisterRoutes(api, userHandler)
	}

	server := httptest.NewServer(r)
	return server, db
}

// TestE2E tests the end-to-end flow of the application.
// Run with: go test -v ./test
func TestE2E(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Load config and override DSN for sqlite in-memory
	cfg := config.Load()
	// Override DSN not used for SQLite
	jwtSecret := cfg.JWTSecret

	server, _ := setupTestServer(jwtSecret)
	defer server.Close()

	client := &http.Client{Timeout: 5 * time.Second}

	// 1. Register a new user
	registerPayload := map[string]string{
		"email":    "e2e@example.com",
		"password": "password123",
	}
	registerBytes, _ := json.Marshal(registerPayload)
	registerReq, _ := http.NewRequest(http.MethodPost, server.URL+"/api/users/register", bytes.NewBuffer(registerBytes))
	registerReq.Header.Set("Content-Type", "application/json")
	registerResp, err := client.Do(registerReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, registerResp.StatusCode)

	// 2. Login with the registered user
	loginPayload := registerPayload
	loginBytes, _ := json.Marshal(loginPayload)
	loginReq, _ := http.NewRequest(http.MethodPost, server.URL+"/api/users/login", bytes.NewBuffer(loginBytes))
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp, err := client.Do(loginReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, loginResp.StatusCode)

	loginBody, _ := ioutil.ReadAll(loginResp.Body)
	var loginResult map[string]string
	err = json.Unmarshal(loginBody, &loginResult)
	assert.NoError(t, err)
	jwtToken, ok := loginResult["token"]
	assert.True(t, ok)
	assert.NotEmpty(t, jwtToken)

	// 3. Use the JWT token to access protected demo endpoint (create a demo)
	demoPayload := map[string]string{
		"name": "Test Demo",
	}
	demoBytes, _ := json.Marshal(demoPayload)
	createDemoReq, _ := http.NewRequest(http.MethodPost, server.URL+"/api/demos/", bytes.NewBuffer(demoBytes))
	createDemoReq.Header.Set("Content-Type", "application/json")
	createDemoReq.Header.Set("Authorization", "Bearer "+jwtToken)
	createDemoResp, err := client.Do(createDemoReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, createDemoResp.StatusCode)

	// 4. Get demos using the JWT token
	getDemoReq, _ := http.NewRequest(http.MethodGet, server.URL+"/api/demos/", nil)
	getDemoReq.Header.Set("Authorization", "Bearer "+jwtToken)
	getDemoResp, err := client.Do(getDemoReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getDemoResp.StatusCode)

	getDemoBody, _ := ioutil.ReadAll(getDemoResp.Body)
	var demos []demo.Demo
	err = json.Unmarshal(getDemoBody, &demos)
	assert.NoError(t, err)
	// We expect at least one demo (the one we just created)
	assert.GreaterOrEqual(t, len(demos), 1)
} 