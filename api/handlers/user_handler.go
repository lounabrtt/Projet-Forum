package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"ff/api/validators"
	"ff/database/models"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		lastName := r.FormValue("lastName")
		firstName := r.FormValue("firstName")
		birthdate := r.FormValue("birthdate")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")

		u2, _ := uuid.NewV4()

		if !validators.ComparePasswords(password, confirmPassword) {
			fmt.Println("Passwords do not match")
			http.Error(w, "Passwords do not match", http.StatusBadRequest)
			return
		}

		var existingUser string
		err := db.QueryRow("SELECT email FROM users WHERE email = ?", email).Scan(&existingUser)

		if err != nil && err != sql.ErrNoRows {
			fmt.Println("Error checking user:", err) //ici
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if existingUser != "" {
			fmt.Println("User already exists")
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		_, err = db.Exec("INSERT INTO users (UUID, lastName, firstName, birthdate, email, password, confirmPassword) VALUES (?, ?, ?, ?, ?, ?, ?)", u2.String(), lastName, firstName, birthdate, email, password, confirmPassword)

		if err != nil {
			fmt.Println("Error inserting user:", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/news", http.StatusSeeOther)
		return

	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var storedHashedPassword string
		err := db.QueryRow("SELECT hashedPassword FROM users WHERE email = ?", email).Scan(&storedHashedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("User not found")
				http.Error(w, "Invalid email or password", http.StatusUnauthorized)
				return
			}
			fmt.Println("Error checking user:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Comparer le mot de passe haché
		err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))
		if err != nil {
			fmt.Println("Invalid password")
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// L'utilisateur est authentifié, gérer la session ou les jetons ici
		// Par exemple, vous pouvez utiliser des cookies de session ou des JWT

		// Exemple de redirection après succès de la connexion
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
