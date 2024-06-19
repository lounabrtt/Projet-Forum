package actions

import (
	"database/sql"
	"ff/api/validators"
	"ff/database/models"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func InitDefaultUsers(db *sql.DB) error {
	users := []models.User{
		{
			UUID:      uuid.Must(uuid.NewV4()).String(),
			LastName:  "Doe",
			FirstName: "Jane",
			Birthdate: "1995-01-01",
			Email:     "admin@beep-bo.com",
			Password:  "123",
			Role:      "admin",
		},
	}

	for _, user := range users {
		exists, err := userExists(db, user.Email)
		if err != nil {
			return fmt.Errorf("failed to check if user exists: %v", err)
		}

		if exists {
			log.Printf("User with email %s already exists. Skipping insertion.", user.Email)
			continue
		}
		if err := CreateDefaultUser(db, user.UUID, user.LastName, user.FirstName, user.Birthdate, user.Email, user.Password, user.Role); err != nil {
			return fmt.Errorf("failed to insert default user: %v", err)
		}
		
	}

	return nil
}

func CreateDefaultUser(db *sql.DB, uuid, lastName, firstName, birthdate, email, password, role string) error {
	hashedPassword, err := validators.HashPassword(password)
	if err != nil {
		log.Printf("Error hashing password for user %s: %v", email, err)
		return fmt.Errorf("error hashing password: %v", err)
	}

	log.Printf("Hashed password for user %s: %s", email, hashedPassword)

	_, err = db.Exec(`
        INSERT INTO users (UUID, lastName, firstName, birthdate, email, password, role)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `, uuid, lastName, firstName, birthdate, email, hashedPassword, role)
	if err != nil {
		log.Printf("Error inserting user %s: %v", email, err)
		return fmt.Errorf("error inserting default user: %v", err)
	}
	log.Printf("Successfully inserted user %s", email)
	return nil
}

func userExists(db *sql.DB, email string) (bool, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM users WHERE email = ?`, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking if user exists: %v", err)
	}
	return count > 0, nil
}