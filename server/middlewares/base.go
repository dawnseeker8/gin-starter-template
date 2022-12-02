package middlewares

import (
	"dawnseek.com/gin-starter/config"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, cfg *config.Config) {
	SetupLogger(r)
	SetupSessions(r, cfg)
}
