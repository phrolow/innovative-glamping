package handlers

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrorHandler middleware to handle errors globally
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{
					Message: "Internal Server Error",
					Code:    http.StatusInternalServerError,
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
