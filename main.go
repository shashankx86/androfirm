package main

import (
	"af/internal"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	// Determine the configuration file path based on the binary name
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf("Could not determine executable path: %s\n", err)
	}
	binaryName := filepath.Base(executable)

	var configPath string
	if strings.HasSuffix(binaryName, ".out") {
		configPath = "./configs/config.yaml"
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Could not determine home directory: %s\n", err)
		}
		configPath = filepath.Join(homeDir, ".af.yaml")
	}

	// Set up configuration
	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file (%s): %s", configPath, err)
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
