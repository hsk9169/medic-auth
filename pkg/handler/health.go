package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// HealthCheckPath is the path to the health check endpoint
	HealthCheckPath = "/health/ready"
)

type HealthCheckHandler struct{}

func (h HealthCheckHandler) GetPathHttpMethod() (string, string, bool) {
	return HealthCheckPath, http.MethodGet, false
}

// Handle health check
// @Summary health check
// @Description health check
// @Produce json
// @Success 200 {string} string "ok"
// @Router /health/ready [get]
func (h HealthCheckHandler) Handle(c *gin.Context) {
	c.Status(http.StatusOK)
}
