package middleware_test

import (
	"innovative_glamping/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	// Generate a valid token
	token, _ := GenerateJWT("test_user")

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{"Valid Token", "Bearer " + token, http.StatusOK},
		{"Missing Token", "", http.StatusUnauthorized},
		{"Invalid Token", "Bearer invalid_token", http.StatusUnauthorized},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/", nil)
		if test.authHeader != "" {
			req.Header.Set("Authorization", test.authHeader)
		}
		rr := httptest.NewRecorder()

		// Wrap a simple handler with the Authenticate middleware
		handler := middleware.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedStatus {
			t.Errorf("%s: got %d, want %d", test.name, rr.Code, test.expectedStatus)
		}
	}
}
