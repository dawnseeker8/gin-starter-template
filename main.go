package main

import (
	"time"

	"dawnseek.com/gin-starter/config"
	"dawnseek.com/gin-starter/handlers"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	config, err := config.LoadConfig()

	if err != nil {
		logger.Debug("cannot load config: ", zap.String("error", err.Error()))
	}

	router := gin.Default()

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	router.Use(ginzap.RecoveryWithZap(logger, true))

	router.GET("/health-check", handlers.HandleHealthCheck)

	router.Run(config.ServerAddress)
}
