package routes

import (
	"github.com/HarvinRaj/goldshop/internal/users"
	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {

	userRepo := users.NewUserRepository()
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(userService)

	handlers := &Handlers{
		UserHandler: userHandler,
	}

	router := gin.Default()
	users.RegisterRoutes(router, handlers.UserHandler)
	return router
}
