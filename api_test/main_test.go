package api_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "localhost:8080/home", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "localhost:8080/home", req.Host)
}
