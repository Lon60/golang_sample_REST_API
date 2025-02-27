package user

import "github.com/gin-gonic/gin"

// RegisterRoutes registers user related routes.
func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	userRoutes := rg.Group("/users")
	{
		userRoutes.POST("/register", handler.Register)
		userRoutes.POST("/login", handler.Login)
	}
} 