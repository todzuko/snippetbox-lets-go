package main

import (
	"github.com/todzuko/snippetbox-lets-go/internal/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "pong")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name         string
		urlPath      string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid ID",
			urlPath:      "/snippet/view/1",
			expectedCode: http.StatusOK,
			expectedBody: "Mock snippet content",
		},
		{
			name:         "Not found ID",
			urlPath:      "/snippet/view/123",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Negative ID",
			urlPath:      "/snippet/view/-1",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Decimal ID",
			urlPath:      "/snippet/view/1.23",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "String ID",
			urlPath:      "/snippet/view/foo",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Empty ID",
			urlPath:      "/snippet/view/",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			assert.Equal(t, tt.expectedCode, code)

			if tt.expectedBody != "" {
				assert.StringContains(t, body, tt.expectedBody)
			}
		})
	}
}

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	validCsrfToken := extractCSRFToken(t, body)

	const (
		validName     = "User"
		validEmail    = "user@example.com"
		validPassword = "password"
		formTag       = "<form action='/user/signup' method='POST' novalidate>"
	)

	tests := []struct {
		name            string
		userName        string
		userEmail       string
		userPassword    string
		csrfToken       string
		expectedCode    int
		expectedFormTag string
	}{
		{
			name:         "Valid name",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    validCsrfToken,
			expectedCode: http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF Token",
			userName:     validName,
			userEmail:    validEmail,
			userPassword: validPassword,
			csrfToken:    "invalidToken",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:            "Empty username",
			userName:        "",
			userEmail:       validEmail,
			userPassword:    validPassword,
			csrfToken:       validCsrfToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Empty email",
			userName:        validName,
			userEmail:       "",
			userPassword:    validPassword,
			csrfToken:       validCsrfToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Invalid email",
			userName:        validName,
			userEmail:       "emiail@test.",
			userPassword:    validPassword,
			csrfToken:       validCsrfToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Empty password",
			userName:        validName,
			userEmail:       validEmail,
			userPassword:    "",
			csrfToken:       validCsrfToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Short password",
			userName:        validName,
			userEmail:       validEmail,
			userPassword:    "pass",
			csrfToken:       validCsrfToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
		{
			name:            "Duplicate email",
			userName:        validName,
			userEmail:       "dupe@test.com",
			userPassword:    validPassword,
			csrfToken:       validCsrfToken,
			expectedCode:    http.StatusUnprocessableEntity,
			expectedFormTag: formTag,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, code, tt.expectedCode)

			if tt.expectedFormTag != "" {
				assert.StringContains(t, body, tt.expectedFormTag)
			}
		})
	}
}
