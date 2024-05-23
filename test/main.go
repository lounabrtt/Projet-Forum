package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"test/functions"

	_ "github.com/mattn/go-sqlite3"
)

const port = ":8080"

func main() {
	//open the database with sqlite3
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		panic(err.Error())
	}
	//creat the 2 tables
	functions.CreateTable(db)
	defer db.Close()

	//handle the different pages
	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) { functions.Home(w, r, db) })
	// http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) { functions.Profile(w, r, db) })
	// http.HandleFunc("/createPost", func(w http.ResponseWriter, r *http.Request) { functions.CreatePost(w, r, db) })
	// http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) { functions.Signin(w, r, db) })
	// http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) { functions.Signup(w, r, db) })
	// http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) { functions.Logout(w, r, db) })

	//load the CSS and the images
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	//start the local host
	fmt.Println("\n(http://localhost:8080/home) - Server started on port", port)
	http.ListenAndServe(port, nil)
}
