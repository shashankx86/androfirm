package internal

import (
	"af/pkg/api"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router and the API routes.
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Apply the rate limit middleware to all routes
	router.Use(api.RateLimitMiddleware)

	// Define the routes
	router.HandleFunc("/api/ping", api.PingHandler).Methods("GET")
	router.HandleFunc("/api/firmware", api.FirmwareHandler).Methods("GET")

	return router
}
