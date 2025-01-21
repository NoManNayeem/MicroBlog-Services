package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Import SQLite driver with a blank identifier
)

var DB *sql.DB // Global database variable

// InitializeDatabase sets up the SQLite3 database connection
func InitializeDatabase() *sql.DB {
	// Connect to SQLite database file
	db, err := sql.Open("sqlite", "comments.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create the comments table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	DB = db // Assign the database connection to the global variable
	return db
}
