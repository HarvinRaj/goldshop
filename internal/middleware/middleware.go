package middleware

import (
	"github.com/HarvinRaj/goldshop/internal/auth"
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func DatabaseMiddleWare(database *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", database)
		c.Next()
	}
}

func AuthMiddleware(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Get token header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing, empty or missing 'Bearer'",
			})

			c.Abort()
			return
		}

		//Validate token with claims
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired claims",
			})

			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
