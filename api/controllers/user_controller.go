package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"ff/database/models"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetAllUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT UUID, lastName, firstName, birthdate, email, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UUID, &user.LastName, &user.FirstName, &user.Birthdate, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateRoleByUserUUID(uuid string, role string) error {
	_, err := db.Exec("UPDATE users SET role = ? WHERE UUID = ?", role, uuid)
	if err != nil {
		return fmt.Errorf("error updating user role: %v", err)
	}
	return nil
}

func UpdateInfoByUserUUID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	currentUser, err := GetCurrentLoggedInUser(r)
	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		log.Printf("Error fetching user: %v", err)
		return
	}

	uuid := currentUser.UUID
	lastName := r.FormValue("lastName")
	firstName := r.FormValue("firstName")
	email := r.FormValue("email")

	_, err = db.Exec("UPDATE users SET lastName = ?, firstName = ?, email = ? WHERE UUID = ?", lastName, firstName, email, uuid)
	if err != nil {
		http.Error(w, "Failed to update user info", http.StatusInternalServerError)
		log.Printf("Error updating user info: %v", err)
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
