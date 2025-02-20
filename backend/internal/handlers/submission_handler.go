package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

type SubmissionHandler struct {
	subService *service.SubmissionService
}

func NewSubmissionHandler(subService *service.SubmissionService) *SubmissionHandler {
	return &SubmissionHandler{
		subService: subService,
	}
}

func (h *SubmissionHandler) CreateSubmission(c *gin.Context) {

	var sub models.Submission
	if err := c.ShouldBindJSON(&sub); err != nil {
		logger.Error("Invalid form data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid submission data"})
		return
	}

	if err := h.subService.CreateSubmission(&sub); err != nil {
		logger.Error("Failed to create form", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Form created successfully", zap.String("formId", sub.ID.Hex()))
	c.JSON(http.StatusCreated, sub)
}
