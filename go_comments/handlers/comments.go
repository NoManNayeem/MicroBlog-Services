package handlers

import (
	"encoding/json"
	"go_comments/database"
	"go_comments/models"
	"log"
	"net/http"
	"os"
)

// HandleComments manages CRUD operations for comments
func HandleComments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getComments(w, r)
	case http.MethodPost:
		createComment(w, r)
	default:
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

// getComments retrieves all comments from the database
func getComments(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, post_id, title, content FROM comments")
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve comments"}`, http.StatusInternalServerError)
		log.Printf("Error retrieving comments: %v", err)
		return
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Title, &comment.Content); err != nil {
			http.Error(w, `{"error": "Failed to parse comment"}`, http.StatusInternalServerError)
			log.Printf("Error parsing comment: %v", err)
			return
		}
		comments = append(comments, comment)
	}

	// Send the comments as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		http.Error(w, `{"error": "Failed to encode comments"}`, http.StatusInternalServerError)
		log.Printf("Error encoding comments: %v", err)
	}
}

// createComment adds a new comment to the database
func createComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Validate the input
	if comment.PostID == 0 || comment.Title == "" || comment.Content == "" {
		http.Error(w, `{"error": "Post ID, title, and content are required"}`, http.StatusBadRequest)
		log.Println("Validation failed: Missing Post ID, title, or content")
		return
	}

	// Validate that the PostID exists in the Flask app
	bearerToken := extractBearerToken(r)
	if !validatePostID(comment.PostID, bearerToken) {
		http.Error(w, `{"error": "Invalid Post ID"}`, http.StatusBadRequest)
		log.Printf("Validation failed: Post ID %d does not exist", comment.PostID)
		return
	}

	// Insert the comment into the database
	_, err := database.DB.Exec("INSERT INTO comments (post_id, title, content) VALUES (?, ?, ?)",
		comment.PostID, comment.Title, comment.Content)
	if err != nil {
		http.Error(w, `{"error": "Failed to save comment"}`, http.StatusInternalServerError)
		log.Printf("Error saving comment to database: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Comment created successfully"}`))
	log.Printf("Comment created successfully for Post ID: %d", comment.PostID)
}

// validatePostID checks if the provided post_id exists in the Flask app
func validatePostID(postID int, bearerToken string) bool {
	flaskURL := os.Getenv("FLASK_APP_URL")
	if flaskURL == "" {
		log.Println("Environment variable FLASK_APP_URL is not set")
		return false
	}

	posts, err := GetPosts(bearerToken)
	if err != nil {
		log.Printf("Error fetching posts from Flask app: %v", err)
		return false
	}

	log.Printf("Fetched posts: %+v", posts) // Log fetched posts for debugging
	for _, post := range posts {
		if post.ID == postID {
			return true
		}
	}
	log.Printf("Post ID %d not found in Flask app", postID)
	return false
}

// extractBearerToken extracts the Bearer token from the Authorization header
func extractBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:]
	}
	return ""
}
