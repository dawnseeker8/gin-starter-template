package handlers

import (
	"github.com/gin-gonic/gin"
)

// HandleHealthCheck handler to health-check re4quest
func HandleHealthCheck(c *gin.Context) {

	c.String(200, "OK")
}
