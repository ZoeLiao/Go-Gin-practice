package main

import (
	"Go-Gin-practice/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_setupRouter(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	test_url := "http://google.com"
	data := url.Values{"path": {"apple"}, "url": {test_url}}
	body := strings.NewReader(data.Encode())
	req, _ := http.NewRequest("POST", "/api/v1/shortener", body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), test_url)
}
