package controllers

import (
	"encoding/json"
	"net/http"
)

func GetSample(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a sample response
	response := map[string]string{
		"message": "Hello from the sample endpoint!",
		"status":  "success",
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
