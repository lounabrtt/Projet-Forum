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
            password VARCHAR(12) NOT NULL,
			confirmPassword VARCHAR(12) NOT NULL,
			isAdmin BOOL NOT NULL DEFAULT FALSE,
			isBanned BOOL NOT NULL DEFAULT FALSE,
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

