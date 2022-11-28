package middlewares

import (
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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
