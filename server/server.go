package server

import (
	"dawnseek.com/gin-starter/config"
	_ "dawnseek.com/gin-starter/docs" // docs is generated by Swag CLI, you have to import it.
	"dawnseek.com/gin-starter/server/middlewares"
	"dawnseek.com/gin-starter/server/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Start() {

	r, cfg := InitServer()

	r.Run(cfg.ServerAddress)
}

func InitServer() (*gin.Engine, config.Config) {
	r := gin.Default()

	cfg := config.LoadConfig()
	config.InitAuthConfig(&cfg)

	middlewares.Init(r, &cfg)
	routers.Init(r, &cfg)

	return r, cfg
}
