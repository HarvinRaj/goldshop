package routes

import (
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/users"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/gin-gonic/gin"
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

	logger.InfoLog.Println("Preparing to start server")

	if err := router.Run(config.Port); err != nil {
		logger.ErrorLog.Printf("Failed to start router, %s", err)
		return err
	}

	return nil
}
