package database

import (
	"database/sql"
	"fmt"
)



func InitTables(db *sql.DB) error {
	if err := CreateTableUser(db); err != nil {
		return fmt.Errorf("error creating user table: %v", err)
	}
	if err := CreateTableCategories(db); err != nil {
		return fmt.Errorf("error creating categories table: %v", err)
	}
	if err := CreateTablePost(db); err != nil {
		return fmt.Errorf("error creating post table: %v", err)
	}
	if err := CreateTableComment(db); err != nil {
		return fmt.Errorf("error creating comments table: %v", err)
	}
	if err := CreateTableNews(db); err != nil {
		return fmt.Errorf("error creating comments table: %v", err)
	}
	return nil
}

// Users
func CreateTableUser(db *sql.DB) error {
	// Creating the user table if not already created
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
			UUID VARCHAR(36) PRIMARY KEY NOT NULL,
            lastName VARCHAR(12) NOT NULL,
			firstName VARCHAR(12) NOT NULL,
			birthdate TEXT NOT NULL,
			email TEXT NOT NULL,
            password VARCHAR(255) NOT NULL,
			role VARCHAR(20) DEFAULT 'user',
			pp BLOB
        )
    `)

	return err
}


// Posts

func CreateTableCategories(db *sql.DB) error {
	// Creating the categories table if not already created
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL UNIQUE,
			number_of_posts INTEGER DEFAULT 0
		)
	`)
	return err
}

func CreateTablePost(db *sql.DB) error {
	// Creating the post table if not already created
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			UUID VARCHAR(36) PRIMARY KEY NOT NULL,
			title VARCHAR(100) NOT NULL,
			content TEXT NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			category TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`)
	return err
}

func CreateTableComment(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			UUID VARCHAR(36) PRIMARY KEY NOT NULL,
			postUUID VARCHAR(36) NOT NULL,
			author VARCHAR(36) NOT NULL,
			content TEXT NOT NULL,
			date TEXT NOT NULL,
			FOREIGN KEY(postUUID) REFERENCES posts(UUID),
			FOREIGN KEY(author) REFERENCES users(UUID)
		)
	`)
	return err
}

func CreateTableNews(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS news (
			UUID VARCHAR(36) PRIMARY KEY NOT NULL,
			title VARCHAR(100) NOT NULL,
			content TEXT NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			category TEXT NOT NULL,
			author VARCHAR(36) NOT NULL
		)
	`)
	return err
}

