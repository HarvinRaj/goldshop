package routes

import (
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/HarvinRaj/goldshop/internal/middleware"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/gin-gonic/gin"
)

// NewRoutes ...
func NewRoutes(config *configs.Config, db *db.DB) error {

	router := gin.Default()
	registerNewRoutes(router, db)

	logger.InfoLog.Info.Println("Preparing to start server")

	if err := router.Run(config.Port); err != nil {
		logger.ErrorLog.Error.Printf("Failed to start router, %v", err)
		return err
	}

	return nil
}

// Register all the routes here
func registerNewRoutes(router *gin.Engine, db *db.DB) {

	router.Use(middleware.DatabaseMiddleWare(db))
	//User Routes
	RegisterUserRoutes(router, db)
	logger.InfoLog.Success.Println("User routes registered")
}
