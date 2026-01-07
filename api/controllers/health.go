package controllers

import (
	"github.com/gin-gonic/gin"
	"gnt-cc/model"
)

type HealthController struct{}

// GetHealth godoc
// @Summary Health check endpoint
// @Description Returns the health status of the gnt-cc API
// @Produce json
// @Success 200 {object} model.HealthResponse
// @Router /health [get]
func (controller *HealthController) GetHealth(c *gin.Context) {
	c.JSON(200, model.HealthResponse{
		Message: "gnt-cc is healthy :)",
	})
}
