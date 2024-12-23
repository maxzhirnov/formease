package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type FormHandler struct {
	formService *service.FormService
}

func NewFormHandler(formService *service.FormService) *FormHandler {
	return &FormHandler{
		formService: formService,
	}
}

func (h *FormHandler) CreateForm(c *gin.Context) {
	// Add debug logging
	logger.Info("Request headers", zap.String("authorization", c.GetHeader("Authorization")))

	var form models.Form
	if err := c.ShouldBindJSON(&form); err != nil {
		logger.Error("Invalid form data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Debug logging for context
	logger.Info("Context values", zap.Any("keys", c.Keys))

	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("UserID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	logger.Info("UserID from context", zap.Any("userID", userID))

	// Convert the userID to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		logger.Error("Invalid user ID format", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	// Set the UserID field of the form
	form.UserID = objectID

	if err := h.formService.CreateForm(&form); err != nil {
		logger.Error("Failed to create form", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Form created successfully", zap.String("formId", form.ID.Hex()))
	c.JSON(http.StatusCreated, form)
}

func (h *FormHandler) GetForm(c *gin.Context) {
	id := c.Param("id")
	logger.Info("Get form", zap.String("formId", id))

	form, err := h.formService.GetForm(id)
	if err != nil {
		logger.Error("Failed to get form", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Form retrieved successfully", zap.String("formId", id))
	c.JSON(http.StatusOK, form)
}

func (h *FormHandler) ListForms(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		logger.Error("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	logger.Info("User ID", zap.String("userID", userID.(string)))

	UserIDString, ok := userID.(string)
	if !ok {
		logger.Error("Invalid user ID type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	forms, err := h.formService.ListForms(UserIDString)
	if err != nil {
		logger.Error("Failed to list forms", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Forms listed successfully", zap.Int("count", len(forms)))
	c.JSON(http.StatusOK, forms)
}

// type CustomForm struct {
// 	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	UserID          string             `bson:"user_id" json:"user_id" binding:"required"`
// 	IsDraft         bool               `bson:"is_draft" json:"is_draft"`
// 	Name            string             `bson:"name" json:"name"`
// 	Theme           string             `bson:"theme" json:"theme"`
// 	FloatingShapes  string             `bson:"floatingShapesTheme" json:"floatingShapesTheme"`
// 	Questions       []Question         `bson:"questions" json:"questions"`
// 	ThankYouMessage ThankYouMessage    `bson:"thankYouMessage" json:"thankYouMessage"`
// }

func (h *FormHandler) UpdateForm(c *gin.Context) {
	id := c.Param("id")
	logger.Info("Update form", zap.String("formId", id))
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Invalid form ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form ID"})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		logger.Error("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var form models.Form
	if err := c.ShouldBindJSON(&form); err != nil {
		logger.Error("Invalid form data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	logger.Info("Form data received", zap.Any("formData", form))

	form.UserID, err = primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		logger.Error("Invalid user ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	form.ID = objectID

	if err := h.formService.UpdateForm(&form); err != nil {
		logger.Error("Failed to update form", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Form updated successfully", zap.String("formId", id))
	c.JSON(http.StatusOK, form)
}

func (h *FormHandler) DeleteForm(c *gin.Context) {
	id := c.Param("id")

	if err := h.formService.DeleteForm(id); err != nil {
		logger.Error("Failed to delete form", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Form deleted successfully", zap.String("formId", id))
	c.JSON(http.StatusOK, gin.H{"message": "Form deleted successfully"})
}
