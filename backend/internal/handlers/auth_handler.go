package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/service"
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

	// Set access token cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"access_token",
		accessToken,
		3600, // 1 hour
		"/",
		"",   // domain
		true, // secure
		true, // httpOnly
	)

	// Set refresh token cookie with specific path
	c.SetCookie(
		"refresh_token",
		refreshToken,
		7*24*3600,  // 7 days
		"/refresh", // restrict to refresh endpoint
		"",         // domain
		true,       // secure
		true,       // httpOnly
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
	// Get refresh token from cookie instead of request body
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "No refresh token provided"})
		return
	}

	accessToken, newRefreshToken, err := h.userService.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		// Clear cookies if refresh token is invalid
		c.SetCookie("refresh_token", "", -1, "/refresh", "", true, true)
		c.SetCookie("access_token", "", -1, "/", "", true, true)
		c.JSON(401, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Set new access token cookie
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"access_token",
		accessToken,
		3600, // 1 hour
		"/",
		"",   // domain
		true, // secure
		true, // httpOnly
	)

	// Set new refresh token cookie
	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		7*24*3600,  // 7 days
		"/refresh", // restrict to refresh endpoint
		"",         // domain
		true,       // secure
		true,       // httpOnly
	)

	// Return success without tokens in body
	c.JSON(200, gin.H{"status": "success"})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userId := c.GetString("userId")
	email := c.GetString("email")
	c.JSON(200, gin.H{"userId": userId, "email": email})
}
