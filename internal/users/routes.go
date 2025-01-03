package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *UserHandler) {

	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/register", h.CreateUser)
		}

	}

}
