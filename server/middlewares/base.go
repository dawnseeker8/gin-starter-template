package middlewares

import (
	"dawnseek.com/gin-starter/config"
	"github.com/gin-gonic/gin"
)

// Init initializes the middlewares
func Init(r *gin.Engine, cfg *config.Config) {
	SetupLogger(r)
	SetupSessions(r, cfg)
}
