package routers

import (
	"dawnseek.com/gin-starter/config"
	"dawnseek.com/gin-starter/server/handlers"
	"dawnseek.com/gin-starter/server/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.Engine, cfg *config.Config) {

	r.GET("/health-check", handlers.HandleHealthCheck)

	auth := r.Group("/api")
	{
		auth.POST("/signin", handlers.Signin)
		auth.POST("/logout", middlewares.SessionProtect, handlers.Logout)
		auth.POST("/getAccount", middlewares.SessionProtect, handlers.GetAccount)
	}

	//For Swagger
	if cfg.SwaggerEnabled {
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
