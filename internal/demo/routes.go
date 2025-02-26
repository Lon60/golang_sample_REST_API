package demo

import (
	"github.com/gin-gonic/gin"
	"golang_sample/internal/demo/handler"
)

func RegisterRoutes(rg *gin.RouterGroup, h *handler.DemoHandler) {
	demoRoutes := rg.Group("/demos")
	{
		demoRoutes.POST("/", h.CreateDemo)
		demoRoutes.GET("/", h.GetAllDemos)
		demoRoutes.GET("/:id", h.GetDemo)
		demoRoutes.PUT("/:id", h.UpdateDemo)
		demoRoutes.DELETE("/:id", h.DeleteDemo)
	}
}
