package main

import (
	demo2 "golang_sample/internal/domain/demo"
	"log"

	"github.com/gin-gonic/gin"
	"golang_sample/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

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
	api := r.Group("/api")
	{
		demo2.RegisterRoutes(api, demoHandler)
	}

	r.Run(cfg.Port)
}
