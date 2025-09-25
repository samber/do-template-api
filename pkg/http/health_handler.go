package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
)

// HealthHandler handles health check requests
// This demonstrates how to implement health check endpoint
type HealthHandler struct {
	logger zerolog.Logger `do:""`
}

// NewHealthHandler creates a new HealthHandler with dependency injection
// This demonstrates how to initialize health handlers with dependencies
func NewHealthHandler(injector do.Injector) (*HealthHandler, error) {
	return do.MustInvokeStruct[*HealthHandler](injector), nil
}

// healthCheck handles health check requests
func (h *HealthHandler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:  "healthy",
		Service: "do-template-api",
	})
}