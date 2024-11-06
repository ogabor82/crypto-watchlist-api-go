package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"crypto-watchlist-api/models"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate email format
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Check if email already exists
	existingUser, err := models.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	// Validate password
	req.Password = strings.TrimSpace(req.Password)
	if req.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters long", http.StatusBadRequest)
		return
	}

	// Create user
	user, err := models.CreateUser(req.Email, req.Password, req.Name)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	response := UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if email and password are not empty
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	user, err := models.VerifyCredentials(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	response := UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
