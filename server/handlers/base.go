package handlers

import (
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// APIHandler is the base handler for all the handlers
type APIHandler struct {
	*gin.Context
}

// GetSessionClaims gets the claims from the session
func (a *APIHandler) GetSessionClaims() *casdoorsdk.Claims {
	session := sessions.Default(a.Context)
	s := session.Get("user")
	if s == nil {
		a.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return nil
	}

	claims := s.(casdoorsdk.Claims)
	return &claims
}

// SetSessionClaims sets the claims to the session
func (a *APIHandler) SetSessionClaims(claims *casdoorsdk.Claims) {
	session := sessions.Default(a.Context)
	if claims == nil {
		session.Delete("user")
		return
	}

	session.Set("user", *claims)
	session.Save()
}

// GetSessionUser gets the user from the session
func (a *APIHandler) GetSessionUser() *casdoorsdk.User {
	claims := a.GetSessionClaims()
	if claims == nil {

		return nil
	}

	return &claims.User
}
