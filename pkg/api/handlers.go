package api

import (
	"encoding/json"
	"net/http"
)

// PingResponse represents the JSON response structure for the ping endpoint.
type PingResponse struct {
	Message string `json:"message"`
}

// PingHandler handles requests to the /api/ping endpoint.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
