package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/service"
)

type SubmissionHandler struct {
	submissionService *service.SubmissionService
}

func NewSubmissionHandler(submissionService *service.SubmissionService) *SubmissionHandler {
	return &SubmissionHandler{submissionService: submissionService}
}

func (h *SubmissionHandler) SubmitForm(c *gin.Context) {
	var submission models.Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.submissionService.SubmitForm(&submission); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Form submitted successfully"})
}
