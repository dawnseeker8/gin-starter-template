package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"dawnseek.com/gin-starter/core/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	router := gin.Default()
	cfg := config.Config{}
	setupRouter(router, &cfg)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health-check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
