package main

import (
	"fmt"
	"log"
	"github.com/kiranaClub/config"
	"github.com/kiranaClub/internal/app"
	"github.com/kiranaClub/internal/database"
)

func main() {
	fmt.Println("Starting Server...")
	// Load config
	cfg := config.LoadConfig()
	// Initialize global database instance
	database.DBConnecting(cfg) // Use *cfg to pass the value, not the pointer
	defer database.CloseDB()   // Close DB when app stops
	// Setup router
	e := app.SetupApp()
	// Start the server
	log.Println("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":" + cfg.PORT))
}
