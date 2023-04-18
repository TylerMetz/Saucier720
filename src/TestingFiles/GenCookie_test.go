package main

import (
	"BackendPkg"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestSetCookie(t *testing.T) {
	// Create a fake request and response
	w := httptest.NewRecorder()

	// Set the active user
	activeUser := BackendPkg.User{
		Password: "Password123",
		UserName: "testUser",
	}

	// Set the expected cookie values
	expectedValue := BackendPkg.ValidateUser(activeUser)
	expectedPath := "/"
	expectedHttpOnly := true
	expectedSameSite := http.SameSiteLaxMode
	expectedDomain := "localhost"

	// Call the SetCookie function to set the cookie

	cookie := &http.Cookie{
		Name:     "sessionID",
		Value:    BackendPkg.ValidateUser(activeUser),
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost",
	}
	http.SetCookie(w, cookie)

	// Check that the cookie is set correctly
	setCookieHeader := w.Header().Get("Set-Cookie")
	cookies := strings.Split(setCookieHeader, "; ")
	var actualValue string
	var actualPath string
	var actualHttpOnly bool
	var actualSameSite http.SameSite
	var actualDomain string
	for _, cookie := range cookies {
		parts := strings.Split(cookie, "=")
		if len(parts) != 2 {
			continue
		}
		switch parts[0] {
		case "sessionID":
			actualValue = parts[1]
		case "Path":
			actualPath = parts[1]
		case "HttpOnly":
			actualHttpOnly = strings.ToLower(parts[1]) == "true"
		case "SameSite":
			actualSameSite = http.SameSiteLaxMode
		case "Domain":
			actualDomain = parts[1]
		}
	}
	if actualValue != expectedValue {
		t.Errorf("Expected cookie value %s, got %s", expectedValue, actualValue)
	}
	if actualPath != expectedPath {
		t.Errorf("Expected cookie path %s, got %s", expectedPath, actualPath)
	}
	if actualHttpOnly != expectedHttpOnly {
		t.Errorf("Expected cookie HttpOnly %t, got %t", expectedHttpOnly, actualHttpOnly)
	}
	if actualSameSite != expectedSameSite {
		t.Errorf("Expected cookie SameSite %v, got %v", expectedSameSite, actualSameSite)
	}
	if actualDomain != expectedDomain {
		t.Errorf("Expected cookie domain %s, got %s", expectedDomain, actualDomain)
	}
}
