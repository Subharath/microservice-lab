package main

import (
	"log"
	"order-service/config"
	"order-service/routes"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup router
	router := routes.SetupRouter(cfg)

	// Start server
	log.Printf("[%s] Starting server on port %s (env: %s, version: %s)",
		cfg.ServiceName, cfg.Port, cfg.Environment, cfg.Version)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
