package handlers

import (
	"net/http"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godocs
// @Summary Health Check
// @Description Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Succuss 200 {object} helper.BaseHttpResponse "Succuss"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working", true, 0))
}
