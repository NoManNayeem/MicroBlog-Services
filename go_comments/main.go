package main

import (
	"fmt"
	"go_comments/database"
	"go_comments/handlers"
	"go_comments/middleware"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get Flask URL and token verification endpoint from environment variables
	flaskURL := os.Getenv("FLASK_APP_URL")
	tokenVerifyURL := os.Getenv("TOKEN_VERIFY_URL")

	if flaskURL == "" || tokenVerifyURL == "" {
		log.Fatal("Required environment variables FLASK_APP_URL or TOKEN_VERIFY_URL are missing")
	}

	// Initialize the database
	db := database.InitializeDatabase()
	defer db.Close()

	// Set up routes
	http.Handle("/comments", middleware.ValidateJWT(http.HandlerFunc(handlers.HandleComments), tokenVerifyURL))

	// Get the server port from the environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
