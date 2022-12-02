package handlers

import (
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	*gin.Context
}

func (a *ApiHandler) GetSessionClaims() *casdoorsdk.Claims {
	session := sessions.Default(a.Context)
	s := session.Get("user")
	if s == nil {
		a.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return nil
	}

	claims := s.(casdoorsdk.Claims)
	return &claims
}

func (a *ApiHandler) SetSessionClaims(claims *casdoorsdk.Claims) {
	session := sessions.Default(a.Context)
	if claims == nil {
		session.Delete("user")
		return
	}

	session.Set("user", *claims)
	session.Save()
}

func (a *ApiHandler) GetSessionUser() *casdoorsdk.User {
	claims := a.GetSessionClaims()
	if claims == nil {

		return nil
	}

	return &claims.User
}
