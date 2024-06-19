package controllers

import (
	"database/sql"
	"ff/api/middlewares"
	"ff/api/validators"
	"ff/database/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

var sessionStore = make(map[string]string)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lastName := r.FormValue("lastName")
	firstName := r.FormValue("firstName")
	birthdate := r.FormValue("birthdate")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")
	role := "user"

	u2, _ := uuid.NewV4()

	if !validators.ComparePasswords(password, confirmPassword) {
		fmt.Println("Passwords do not match")
		middlewares.SendErrorToClient(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	var existingUser string
	err := db.QueryRow("SELECT email FROM users WHERE email = ?", email).Scan(&existingUser)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Error checking user:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if existingUser != "" {
		fmt.Println("User already exists")
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := validators.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (UUID, lastName, firstName, role, birthdate, email, password) VALUES (?, ?, ?,?, ?, ?, ?)", u2.String(), lastName, firstName, role, birthdate, email, hashedPassword)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	SELECT_USER_BY_EMAIL := "SELECT UUID, lastName, firstName, role, birthdate, email, password FROM users WHERE email = ?"

	var currentUser models.User
	var hashedPassword string
	err := db.QueryRow(SELECT_USER_BY_EMAIL, email).Scan(&currentUser.UUID, &currentUser.LastName, &currentUser.FirstName, &currentUser.Role, &currentUser.Birthdate, &currentUser.Email, &hashedPassword)

	if err == sql.ErrNoRows {
		middlewares.SendErrorToClient(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err != nil {
		middlewares.SendErrorToClient(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !validators.CheckPasswordHash(password, hashedPassword) {
		middlewares.SendErrorToClient(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	sessionToken, _ := uuid.NewV4()

	sessionCookie := http.Cookie{}
	sessionCookie.Name = "session_token"
	sessionCookie.Value = sessionToken.String()
	sessionCookie.Expires = time.Now().Add(24 * time.Hour)
	sessionCookie.Secure = false
	sessionCookie.HttpOnly = true
	sessionCookie.Path = "/"

	http.SetCookie(w, &sessionCookie)

	sessionStore[sessionToken.String()] = currentUser.UUID

	http.Redirect(w, r, "/news", http.StatusSeeOther)
}

func GetCurrentLoggedInUser(r *http.Request) (models.User, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		log.Printf("Err: %v", err)
		return models.User{}, err

	}

	userUUID := sessionStore[cookie.Value]

	if userUUID == "" {
		return models.User{}, http.ErrNoCookie
	}

	SELECT_USER_BY_UUID := "SELECT UUID, lastName, firstName, role, birthdate, email FROM users WHERE UUID = ?"

	var currentUser models.User
	err = db.QueryRow(SELECT_USER_BY_UUID, userUUID).Scan(&currentUser.UUID, &currentUser.LastName, &currentUser.FirstName, &currentUser.Role, &currentUser.Birthdate, &currentUser.Email)
	if err != nil {
		return models.User{}, err
	}

	return currentUser, nil
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	middlewares.SendSuccessResponse(w, "Logged out successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
