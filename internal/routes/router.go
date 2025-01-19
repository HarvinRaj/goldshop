package routes

import (
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/users"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes(config *configs.Config) error {

	router := gin.Default()
	registerNewRoutes(router)

	logger.InfoLog.Info.Println("Preparing to start server")

	if err := router.Run(config.Port); err != nil {
		logger.ErrorLog.Printf("Failed to start router, %v", err)
		return err
	}

	return nil
}

// Register all the routes here
func registerNewRoutes(router *gin.Engine) {
	//User Routes
	users.RegisterRoutes(router)
	logger.InfoLog.Success.Println("User routes registered")
}
