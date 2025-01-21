package models

// Comment represents a single blog comment stored in the Go app
type Comment struct {
	ID      int    `json:"id"`      // Unique ID of the comment
	PostID  int    `json:"post_id"` // Associated Post ID from the Flask app
	Title   string `json:"title"`   // Comment title
	Content string `json:"content"` // Comment content
}
