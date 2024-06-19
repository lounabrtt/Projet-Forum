package api

import (
	"ff/api/controllers"
	"net/http"
)

func Routes() {
	// AUTHENTICATION
	http.HandleFunc("/api/auth/signup", controllers.CreateUser)
	http.HandleFunc("/api/auth/login", controllers.LoginUser)
	http.HandleFunc("/api/auth/logout", controllers.LogoutUser)

	// POSTS
	http.HandleFunc("/api/posts", controllers.CreatePost)
	http.HandleFunc("/api/news", controllers.CreateNew)
	http.HandleFunc("/api/posts/comment/", controllers.AddComment)

	// ADMIN
	http.HandleFunc("/api/users", controllers.GetAllUsers)
}
