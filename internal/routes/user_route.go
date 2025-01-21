package routes

import (
	"github.com/HarvinRaj/goldshop/internal/auth"
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/HarvinRaj/goldshop/internal/handlers"
	"github.com/HarvinRaj/goldshop/internal/middleware"
	"github.com/HarvinRaj/goldshop/internal/repositories"
	"github.com/HarvinRaj/goldshop/internal/services"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes ...
func RegisterUserRoutes(router *gin.Engine, db *db.DB, jwtManager *auth.JWTManager) {

	repo := repositories.NewUserRepository(db.GetConnection())
	service := services.NewUserService(repo, jwtManager)
	handler := handlers.NewUserHandler(service)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", handler.RegisterUser)
		v1.POST("/login", handler.Login)

		userGroup := v1.Group("/users")
		userGroup.Use(middleware.AuthMiddleware(jwtManager))
		{
			userGroup.GET("/list", handler.GetUsersList)
		}

	}

}
