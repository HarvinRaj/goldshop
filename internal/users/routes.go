package users

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes ...
func RegisterRoutes(router *gin.Engine) {

	repo := NewUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/register", handler.CreateUser)
		}

	}

}
