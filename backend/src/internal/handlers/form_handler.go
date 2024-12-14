package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/service"
)

type FormHandler struct {
	formService *service.FormService
}

func NewFormHandler(formService *service.FormService) *FormHandler {
	return &FormHandler{formService: formService}
}

func (h *FormHandler) GetForm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form ID"})
		return
	}

	form, err := h.formService.GetForm(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, form)
}
