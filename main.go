package main

import (
	"crypto-watchlist-api/db"
	"crypto-watchlist-api/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Crypto Watchlist API Server Started")

	// Initialize database
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	fmt.Println("Database initialized successfully")

	// Setup routes
	routes.SetupUserRoutes()
	routes.SetupSampleRoutes()

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
