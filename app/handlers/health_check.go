package handlers

import (
	"go-api/helpers"

	"github.com/gin-gonic/gin"
)

func (h *APIGOHandler) Check(c *gin.Context) {
	healthCheck := h.APIGOUsecase.HealthCheck()
	helpers.RespondJSON(c, 200, healthCheck, nil)
}
