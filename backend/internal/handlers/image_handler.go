// handlers/upload.go
package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

type ImageHandler struct {
	imageService *service.ImageService
}

func NewImageHandler(imageService *service.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

func (h *ImageHandler) UploadImage(c *gin.Context) {
	logger.Info("Starting image upload")

	// Получаем ID пользователя из контекста
	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("UserID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем файл из запроса
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		logger.Error("Failed to get file from request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	defer file.Close()

	// Создаем модель изображения
	image := &models.Image{
		UserID:    userID.(string),
		Filename:  header.Filename,
		Size:      header.Size,
		Type:      header.Header.Get("Content-Type"),
		CreatedAt: time.Now(),
	}

	// Загружаем изображение через сервис
	uploadedImage, err := h.imageService.UploadImage(image, file)
	if err != nil {
		logger.Error("Failed to upload image", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Image uploaded successfully",
		zap.String("imageId", uploadedImage.ID.Hex()),
		zap.String("userId", userID.(string)))

	c.JSON(http.StatusCreated, uploadedImage)
}

func (s *ImageHandler) GetUserImages(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		logger.Error("UserID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	page := c.Query("page")
	if page == "" {
		page = "0"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		logger.Error("Invalid page query parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page query parameter"})
		return
	}

	limit := c.Query("limit")
	if limit == "" {
		limit = "100"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		logger.Error("Invalid limit query parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit query parameter"})
		return
	}

	userIDString, ok := userID.(string)
	if !ok {
		logger.Error("Invalid user ID type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	logger.Info("Finding images by user ID", zap.String("userID", userIDString))
	images, err := s.imageService.FindByUserID(userIDString, pageInt, limitInt)
	if err != nil {
		logger.Error("Failed to find images", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalCount, err := s.imageService.CountByUserID(userIDString)
	if err != nil {
		logger.Error("Failed to count user images", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.ImageResponse{
		Images: images,
		Total:  totalCount,
		Page:   pageInt,
		Limit:  limitInt,
	}

	logger.Info("User images retrieved successfully",
		zap.Int("count", len(images)),
		zap.Int64("total", totalCount))
	c.JSON(http.StatusOK, response)
}

func (s *ImageHandler) DeleteImage(c *gin.Context) {
	imageID := c.Param("id")

	userID, ok := c.Get("userID")
	if !ok {
		logger.Error("UserID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDString, ok := userID.(string)
	if !ok {
		logger.Error("Invalid user ID type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Optional: Check if image belongs to user before deleting
	image, err := s.imageService.FindByID(imageID)
	if err != nil {
		logger.Error("Failed to find image", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	if image.UserID != userIDString {
		logger.Error("Unauthorized image deletion attempt",
			zap.String("requestUserID", userIDString),
			zap.String("imageUserID", image.UserID))
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to delete this image"})
		return
	}

	if err := s.imageService.Delete(imageID); err != nil {
		logger.Error("Failed to delete image", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Image deleted successfully",
		zap.String("imageID", imageID),
		zap.String("userID", userIDString))

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
