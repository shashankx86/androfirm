package api

import (
	"encoding/json"
	"net/http"
)

// HelloResponse represents the JSON response structure.
type HelloResponse struct {
	Message string `json:"message"`
}

// HelloHandler is a basic handler that returns a hello message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
