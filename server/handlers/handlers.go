package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary  HandleHealthCheck handler to health-check re4quest
// @Schemes
// @Description health-check end point to return the healthy status of API.
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /health-check [get]
func HandleHealthCheck(c *gin.Context) {
	a := ApiHandler{c}
	a.String(200, "OK")
}
