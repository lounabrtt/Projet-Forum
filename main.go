package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"ff/api"
	"ff/api/controllers"
	"ff/client"
	"ff/database"

	_ "github.com/mattn/go-sqlite3"
)

const port = ":8080"

var db *sql.DB

func main() {
	db, _ = sql.Open("sqlite3", "./database/forum.db")

	defer db.Close()

	if err := database.InitTables(db); err != nil {
		fmt.Println("Error initializing tables:", err)
		return
	}

	controllers.SetDB(db)

	client.Routes()
	api.Routes()

	// Start the server
	fmt.Println("\n(http://localhost:8080) - Server started on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
