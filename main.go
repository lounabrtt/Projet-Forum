package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"ff/api"
	"ff/api/handlers"
	"ff/database"
	"ff/web"

	_ "github.com/mattn/go-sqlite3"
)

// Define the port variable
const port = ":8080"
var db *sql.DB


func main() {
	// Database connection
	db, _ = sql.Open("sqlite3", "./database/forum.db")

	defer db.Close() 

	if err := database.InitTables(db); err != nil {
        fmt.Println("Error initializing tables:", err)
        return
    }

	handlers.SetDB(db)

	web.Handler()
	api.Routes()

	// Start the server
	fmt.Println("\n(http://localhost:8080) - Server started on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}