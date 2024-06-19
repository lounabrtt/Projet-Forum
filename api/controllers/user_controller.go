package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"ff/database/models"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT UUID, lastName, firstName, birthdate, email FROM users")
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UUID, &user.LastName, &user.FirstName, &user.Birthdate, &user.Email)
		if err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
