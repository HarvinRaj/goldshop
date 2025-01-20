package routes

import (
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/HarvinRaj/goldshop/internal/handlers"
	"github.com/HarvinRaj/goldshop/internal/repositories"
	"github.com/HarvinRaj/goldshop/internal/services"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes ...
func RegisterUserRoutes(router *gin.Engine, db *db.DB) {

	repo := repositories.NewUserRepository(db.GetConnection())
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/register", handler.RegisterUser)
			userGroup.GET("/list", handler.GetUsersList)
			userGroup.POST("/login", handler.Login)
		}

	}

}
