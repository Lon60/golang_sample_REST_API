package main

import (
	"log"

	"golang_sample/internal/demo"
	"golang_sample/internal/demo/handler"
	"golang_sample/internal/demo/model"
	"golang_sample/internal/demo/repository"
	"golang_sample/internal/demo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=demo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&model.Demo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	demoRepo := repository.NewDemoRepository(db)
	demoService := service.NewDemoService(demoRepo)
	demoHandler := handler.NewDemoHandler(demoService)

	r := gin.Default()

	api := r.Group("/api")
	{
		demo.RegisterRoutes(api, demoHandler)
	}

	r.Run(":8080")
}
