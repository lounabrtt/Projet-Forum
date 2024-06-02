package function

import (
	"database/sql"
	// "fmt"
	// "net/http"
)

func CreateTablePost(db *sql.DB) error {
	// Creating the post table if not already created
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(100) NOT NULL,
			content TEXT NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			category INTEGER NOT NULL,
			author INTEGER NOT NULL,
			FOREIGN KEY(category) REFERENCES categories(id),
			FOREIGN KEY(author) REFERENCES users(id)
		)
	`)
	return err
}

func newPost(db *sql.DB, title, content, username string) error {

    if err != nil {
        return err
    }

    // Inserting the post into the database
    _, err = db.Exec(`
        INSERT INTO posts (title, content, username) VALUES (?, ?, ?)
    `, title, content, username)
    return err
}