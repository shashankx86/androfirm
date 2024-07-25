package main

import (
	"af/internal"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Set up configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Get the port from the config file
	port := viper.GetString("server.port")
	if port == "" {
		log.Fatal("Server port must be specified in the config file")
	}

	// Set up router
	router := internal.SetupRouter()

	// Start server
	log.Printf("Starting server on :%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
