package main

import (
	"github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
    "net/url"
	"testing"
    "strings"
)

func Test_setupRouter(t *testing.T) {
    router := gin.Default()
	w := httptest.NewRecorder()
    test_url := "http://google.com"
    data := url.Values{"URL": test_url}
    body := strings.NewReader(data.Encode())
    req, _ := http.NewRequest("POST", "/api/v1/shortener", body)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), test_url)
}
