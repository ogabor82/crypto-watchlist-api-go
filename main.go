package main

import (
	"fmt"
	"net/http"
)

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	fmt.Println("Crypto Watchlist API Server Started")

	// Register the /sample endpoint
	http.HandleFunc("/sample", sampleHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
