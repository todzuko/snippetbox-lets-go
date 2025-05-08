package main

import (
	"bytes"
	"github.com/todzuko/snippetbox-lets-go/internal/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommonHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	commonHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	expected := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, rs.Header.Get("Content-Security-Policy"), expected)

	expected = "origin-when-cross-origin"
	assert.Equal(t, rs.Header.Get("Referrer-Policy"), expected)

	expected = "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expected)

	expected = "deny"
	assert.Equal(t, rs.Header.Get("X-Frame-Options"), expected)

	expected = "0"
	assert.Equal(t, rs.Header.Get("X-XSS-Protection"), expected)

	expected = "Go"
	assert.Equal(t, rs.Header.Get("Server"), expected)

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
