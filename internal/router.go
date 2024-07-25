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

	router.HandleFunc("/api/hello", api.HelloHandler).Methods("GET")
	return router
}
