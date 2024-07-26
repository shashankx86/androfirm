package api

import (
	"af/pkg/scraper"
	"encoding/json"
	"net/http"
)

// PingResponse represents the JSON response structure for the ping endpoint.
type PingResponse struct {
	Message string `json:"message"`
}

// FirmwareResponse represents the JSON response structure for the firmware endpoint.
type FirmwareResponse struct {
	Data string `json:"data"`
}

// PingHandler handles requests to the /api/ping endpoint.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// FirmwareHandler handles requests to the /api/firmware endpoint.
func FirmwareHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the scraper (this could be based on config or other logic)
	scraper := scraper.NewExampleSiteScraper()

	// Fetch data from the scraper
	data, err := scraper.FetchData()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	// Send the response
	response := FirmwareResponse{Data: data}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
