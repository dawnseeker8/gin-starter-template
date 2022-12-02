package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitServer(t *testing.T) {
	router, _ := InitServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health-check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
