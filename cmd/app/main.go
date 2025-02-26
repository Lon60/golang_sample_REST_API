package main

import (
	"github.com/gin-gonic/gin"
	_ "golang_sample/docs"
	"golang_sample/internal/config"
	demo2 "golang_sample/internal/domain/demo"
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

	if err := db.AutoMigrate(&demo2.Demo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	demoRepo := demo2.NewDemoRepository(db)
	demoService := demo2.NewDemoService(demoRepo.Repository)
	demoHandler := demo2.NewDemoHandler(demoService)

	r := gin.Default()

	if cfg.Mode == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	api := r.Group("/api")
	{
		demo2.RegisterRoutes(api, demoHandler)
	}

	r.Run(cfg.Port)
}
