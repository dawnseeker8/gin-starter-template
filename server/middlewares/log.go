package middlewares

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupLogger(r *gin.Engine) {
	logger, _ := zap.NewProduction()

	gin.ForceConsoleColor()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	r.Use(ginzap.RecoveryWithZap(logger, true))
}
