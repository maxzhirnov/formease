package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.Register(c.Request.Context(), &user); err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(201, gin.H{"message": "User registered successfully"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, user, err := h.userService.Login(c.Request.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	logger.Info(fmt.Sprintf("accessToken: %s, refreshToken: %s", accessToken, refreshToken))

	// Set access token cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"access_token",
		accessToken,
		8*3600, // 8 hours
		"/",
		"",
		true,
		true,
	)

	// Set refresh token cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"refresh_token",
		refreshToken,
		7*24*3600,
		"/", // Changed from "/refresh" to "/"
		"",
		true,
		true,
	)

	// Return only user info in response
	c.JSON(200, gin.H{
		"user": gin.H{
			"id":    user.ID.Hex(),
			"email": user.Email,
		},
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	logger.Info("Refresh request received",
		zap.String("path", c.Request.URL.Path),
		zap.Any("headers", c.Request.Header),
		zap.Any("cookies", c.Request.Cookies()))

	// Try to get token from cookie first
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		// If not in cookie, try to get from Cookie header
		cookieHeader := c.GetHeader("Cookie")
		if cookieHeader != "" {
			cookies := strings.Split(cookieHeader, ";")
			for _, cookie := range cookies {
				parts := strings.Split(strings.TrimSpace(cookie), "=")
				if len(parts) == 2 && parts[0] == "refresh_token" {
					refreshToken = parts[1]
					break
				}
			}
		}
	}

	if refreshToken == "" {
		logger.Error("No refresh token found in request")
		c.JSON(401, gin.H{"error": "No refresh token provided"})
		return
	}

	logger.Info("Found refresh token", zap.String("token", refreshToken))

	// Rest of your refresh logic
	accessToken, newRefreshToken, err := h.userService.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		logger.Error("Error refreshing token", zap.Error(err))
		c.SetCookie("refresh_token", "", -1, "/", "", true, true)
		c.SetCookie("access_token", "", -1, "/", "", true, true)
		c.JSON(401, gin.H{"error": "Invalid refresh token"})
		return
	}

	logger.Info("Setting new cookies",
		zap.String("accessToken", accessToken),
		zap.String("refreshToken", newRefreshToken))

	// Set access token cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"access_token",
		accessToken,
		8*3600, // 8 hours
		"/",
		"",
		true,
		true,
	)

	// Set refresh token cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		7*24*3600,
		"/",
		"",
		true,
		true,
	)

	// Add both tokens to response headers for debugging
	c.Header("X-Debug-Access-Token", "set")
	c.Header("X-Debug-Refresh-Token", "set")

	c.JSON(200, gin.H{"status": "success"})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userId := c.GetString("userId")
	email := c.GetString("email")
	c.JSON(200, gin.H{"userId": userId, "email": email})
}
