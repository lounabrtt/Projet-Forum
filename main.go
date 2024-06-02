package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// Define the port variable
const port = ":8080"
var db *sql.DB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func newsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/News.html")
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/login.html")
}
func testlikeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/testlike.html")
}
func testpopupHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/testpopup.html")
}
func adminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/admin.html")
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bjr"))
}
func StaticFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	// Database connection
	db, _ = sql.Open("sqlite3", "./database/forum.db")

	defer db.Close()

	// Create the tables
	if err := CreateTableUser(db); err != nil {
		fmt.Println("Error creating user table:", err)
		return
	}
	if err := CreateTableCategories(db); err != nil {
		fmt.Println("Error creating categories table:", err)
		return
	}
	if err := CreateTablePost(db); err != nil {
		fmt.Println("Error creating post table:", err)
		return
	}

	// HTTP Handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", newsHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/home", indexHandler)
	http.HandleFunc("/css/", StaticFiles)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", createUser)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/testlike", testlikeHandler)
	http.HandleFunc("/testpopup", testpopupHandler)
	// Start the server
	fmt.Println("\n(http://localhost:8080/home) - Server started on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func CreateTableUser(db *sql.DB) error {
	// Creating the user table if not already created
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            nom VARCHAR(12) NOT NULL,
			prenom VARCHAR(12) NOT NULL,
			date TEXT NOT NULL,
			email TEXT NOT NULL,
            password VARCHAR(12) NOT NULL,
			confirmPassword VARCHAR(12) NOT NULL,
			isAdmin BOOL NOT NULL DEFAULT FALSE,
			isBanned BOOL NOT NULL DEFAULT FALSE,
			pp BLOB,
			UUID VARCHAR(36) PRIMARY KEY NOT NULL
        )
    `)
	return err
}
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
func comparePasswords(password, confirmPassword string) bool {

	return password == confirmPassword
}
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		nom := r.FormValue("nom")
		prenom := r.FormValue("prenom")
		date := r.FormValue("date")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")

		u2, _ := uuid.NewV4()
		if comparePasswords(password, confirmPassword) {
			// Check if user already exists
			var existingUser string
			err := db.QueryRow("SELECT email FROM users WHERE email = ?", email).Scan(&existingUser)

			if err != nil && err != sql.ErrNoRows {
				fmt.Println("Error checking user:", err)
				return
			}

			if existingUser != "" {
				fmt.Println("User already exists")
				http.Error(w, "User already exists", http.StatusConflict)
				return
			}

			_, err = db.Exec("INSERT INTO users (nom, prenom, date, email, password, confirmPassword, UUID) VALUES (?, ?, ?, ?, ?, ?, ?)", nom, prenom, date, email, password, confirmPassword, u2.String())

			if err != nil {
				fmt.Println("Error inserting user:", err)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.ServeFile(w, r, "./html/form2.html")

}
