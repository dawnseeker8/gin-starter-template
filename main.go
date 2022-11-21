package main

import (
	"time"

	"dawnseek.com/gin-starter/core/config"
	"dawnseek.com/gin-starter/core/handlers"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	router := setupRouter()

	cfg := config.LoadConfig()

	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	router.Use(ginzap.RecoveryWithZap(logger, true))

	router.Run(cfg.ServerAddress)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health-check", handlers.HandleHealthCheck)
	return r
}
