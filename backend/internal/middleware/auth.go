// middleware/auth.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/utils"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

func AuthMiddleware(jwtUtil *utils.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from cookie instead of header
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			logger.Error("Failed to get access token from cookie", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No access token provided"})
			c.Abort()
			return
		}

		// Validate token
		userId, email, err := jwtUtil.ValidateToken(accessToken)
		if err != nil {
			logger.Error("Failed to validate access token", zap.Error(err))
			// Clear invalid cookie
			c.SetCookie("access_token", "", -1, "/", "", true, true)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("userID", userId)
		c.Set("email", email)

		c.Next()
	}
}
