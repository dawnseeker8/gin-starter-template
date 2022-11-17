package handlers

import (
	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "System running OK",
	})

}
