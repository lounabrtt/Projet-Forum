package actions

import (
	"database/sql"
	"fmt"
)

func InitDefaultCategories(db *sql.DB) error {
	categories := []string{"pop", "rap", "country", "rock", "jazz"}

	for _, category := range categories {
		_, err := db.Exec("INSERT OR IGNORE INTO categories (name) VALUES (?)", category)
		if err != nil {
			return fmt.Errorf("error inserting category %s: %v", category, err)
		}
	}

	return nil
}
