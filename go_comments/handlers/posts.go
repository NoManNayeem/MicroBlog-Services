package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// GetPosts retrieves posts from the Flask app
func GetPosts(bearerToken string) ([]struct {
	ID    int    `json:"id"`    // Post ID
	Title string `json:"title"` // Post title
	Body  string `json:"body"`  // Post content
}, error) {
	// Retrieve Flask URL from the environment
	flaskURL := os.Getenv("FLASK_APP_URL")
	if flaskURL == "" {
		log.Println("Environment variable FLASK_APP_URL is not set")
		return nil, fmt.Errorf("flask URL is not set")
	}

	// Build the Flask blogs endpoint URL
	url := fmt.Sprintf("%s/blogs", flaskURL)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request to Flask app: %v", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the Authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	// Send the GET request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to Flask app: %v", err)
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch posts, HTTP status: %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch posts, HTTP status: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body from Flask app: %v", err)
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response into a slice of Post structs
	var posts []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := json.Unmarshal(body, &posts); err != nil {
		log.Printf("Error unmarshalling posts JSON: %v", err)
		return nil, fmt.Errorf("failed to unmarshal posts: %w", err)
	}

	return posts, nil
}
