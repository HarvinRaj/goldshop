package routes

import (
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/users"
	"github.com/gin-gonic/gin"
	"log"
)

func NewRoutes(config *configs.Config) error {

	userRepo := users.NewUserRepository()
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(userService)

	handlers := &Handlers{
		UserHandler: userHandler,
	}

	router := gin.Default()
	users.RegisterRoutes(router, handlers.UserHandler)

	if err := router.Run(config.Port); err != nil {
		log.Fatalf("Failed to run server, %v", err)
		return err
	}

	log.Printf("Server has started")
	return nil
}
