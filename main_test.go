package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	resp := w.Result()

	// ステータスコード200が返ること
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
