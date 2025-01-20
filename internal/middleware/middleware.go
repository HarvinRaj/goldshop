package middleware

import (
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/gin-gonic/gin"
)

func DatabaseMiddleWare(database *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", database)
		c.Next()
	}
}
