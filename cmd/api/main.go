package main

import (
	"Gin/internal/platform"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env file in the current directory
	env := godotenv.Load()
	if env != nil {
		// It's not an error if the .env file doesn't exist, it may be configured in the production environment.
		// But in development, it's good to warn about it.
		log.Printf("Warning: Error loading .env file (it might not exist if running in production directly): %v", env)
	}

	// Initialize database
	db, err := platform.InitDB()

	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	// Close the database connection when exiting the program
	defer db.Close()

	// Initialize the Gin server
	r := platform.InitGinServer(db)

	// Start the server. Listen on 0.0.0.0:3000
	log.Fatal(r.Run(":3000"))
}
