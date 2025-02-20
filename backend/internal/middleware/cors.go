package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		path := c.Request.URL.Path

		// More permissive for uploads and localhost
		if strings.HasPrefix(path, "/uploads/") ||
			origin == "" ||
			strings.HasPrefix(origin, "http://localhost:") {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		} else {
			// Existing origin checking for other routes
			allowedOrigins := []string{
				"http://localhost:5170",
				"http://localhost:8080",
			}

			isOriginAllowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					isOriginAllowed = true
					c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
					break
				}
			}

			if !isOriginAllowed {
				logger.Error("Unauthorized origin",
					zap.String("origin", origin),
					zap.String("path", path))
				c.AbortWithStatus(403)
				return
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With, Cookie")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
