package main

import (
	"log"
	"net/http"
	"saferoute/internal/api"
	"saferoute/pkg/config"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize router
	router := api.NewRouter()

	// Start server
	log.Printf("Server running on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
