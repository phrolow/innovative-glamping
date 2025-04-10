package handlers

import (
	"encoding/json"
	"net/http"

	"innovative-glamping/middleware"
)

// LoginRequest represents a login request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handles user login and token generation
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	// Simple username/password check (replace with DB validation)
	if req.Username == "admin" && req.Password == "password" {
		token, _ := middleware.GenerateJWT(req.Username)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
