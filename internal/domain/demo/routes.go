package demo

import (
	"github.com/gin-gonic/gin"
	"golang_sample/internal/middleware"
)

func RegisterRoutes(rg *gin.RouterGroup, h *Handler, jwtSecret string) {
	demoRoutes := rg.Group("/demos")
	
	demoRoutes.Use(middleware.JWTAuthMiddleware(jwtSecret))
	{
		demoRoutes.POST("/", h.CreateDemo)
		demoRoutes.GET("/", h.GetAllDemos)
		demoRoutes.GET("/:id", h.GetDemo)
		demoRoutes.PUT("/:id", h.UpdateDemo)
		demoRoutes.DELETE("/:id", h.DeleteDemo)
	}
}
