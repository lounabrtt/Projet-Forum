package api

import (
	"ff/api/handlers"
	"net/http"
)


func Routes() {
	http.HandleFunc("/api/signup", handlers.CreateUser)
	http.HandleFunc("/api/users", handlers.GetAllUsers)
	http.HandleFunc("/api/posts", handlers.CreatePost)
}