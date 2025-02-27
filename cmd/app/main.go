package main

import (
	"github.com/gin-gonic/gin"
	_ "golang_sample/docs"
	"golang_sample/internal/config"
	"golang_sample/internal/domain/demo"
	"golang_sample/internal/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Demo API
// @version 1.0
// @description This is a sample server for a demo API.
// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.Load()

	if cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&demo.Demo{}, &user.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	demoRepo := demo.NewDemoRepository(db)
	demoService := demo.NewDemoService(demoRepo.Repository)
	demoHandler := demo.NewDemoHandler(demoService)

	r := gin.Default()

	if cfg.Mode == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	api := r.Group("/api")
	{
		demo.RegisterRoutes(api, demoHandler, cfg.JWTSecret)

		userRepo := user.NewRepository(db)
		userHandler := user.NewUserHandler(userRepo, cfg.JWTSecret)
		user.RegisterRoutes(api, userHandler)
	}

	r.Run(cfg.Port)
}
