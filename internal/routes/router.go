package routes

import (
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/auth"
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/HarvinRaj/goldshop/internal/middleware"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/gin-gonic/gin"
	"time"
)

// NewRoutes ...
func NewRoutes(config *configs.Config, db *db.DB) error {

	jwtManager := auth.NewJWT(config.SecretKey, time.Hour*1)

	router := gin.Default()
	registerNewRoutes(router, db, jwtManager)

	logger.InfoLog.Info.Println("Preparing to start server")

	if err := router.Run(config.Port); err != nil {
		logger.ErrorLog.Error.Printf("Failed to start router, %v", err)
		return err
	}

	return nil
}

// Register all the routes here
func registerNewRoutes(router *gin.Engine, db *db.DB, jwtManager *auth.JWTManager) {

	router.Use(middleware.DatabaseMiddleWare(db))
	//User Routes
	RegisterUserRoutes(router, db, jwtManager)
	logger.InfoLog.Success.Println("User routes registered")
}
