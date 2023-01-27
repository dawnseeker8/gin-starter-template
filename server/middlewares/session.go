package middlewares

import (
	"net/http"

	"dawnseek.com/gin-starter/config"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// SessionProtect protects the routes with session
func SessionProtect(c *gin.Context) {
	session := sessions.Default(c)
	s := session.Get("claim")
	if s == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims := s.(casdoorsdk.Claims)
	c.Set("claim", claims)
	c.Next()
}

// SetupSessions sets up the session
func SetupSessions(r *gin.Engine, cfg *config.Config) {
	store, _ := redis.NewStore(10, "tcp", cfg.RedisURL, "", []byte(cfg.SessionSecret))
	r.Use(sessions.Sessions(cfg.SessionName, store))
}
