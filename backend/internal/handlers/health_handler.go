// internal/handlers/health_handler.go
package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type HealthHandler struct {
	client *mongo.Client
}

func NewHealthHandler(client *mongo.Client) *HealthHandler {
	return &HealthHandler{client: client}
}

func (h *HealthHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"statusss": "ok",
		"time":     time.Now(),
	})
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	status := "ok"
	dbStatus := "up"

	// Check database connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := h.client.Ping(ctx, readpref.Primary()); err != nil {
		status = "degraded"
		dbStatus = "down"
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"time":   time.Now(),
		"checks": map[string]string{
			"database": dbStatus,
		},
	})
}
