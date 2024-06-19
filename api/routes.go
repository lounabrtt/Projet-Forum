package api

import (
	"ff/api/controllers"
	"log"
	"net/http"
)

func Routes() {
	// AUTHENTICATION
	http.HandleFunc("/api/auth/signup", controllers.CreateUser)
	http.HandleFunc("/api/auth/login", controllers.LoginUser)
	http.HandleFunc("/api/auth/logout", controllers.LogoutUser)

	// USERS
	http.HandleFunc("/api/me/update", controllers.UpdateInfoByUserUUID)

	// POSTS
	http.HandleFunc("/api/posts", controllers.CreatePost)
	http.HandleFunc("/api/news", controllers.CreateNew)
	http.HandleFunc("/api/posts/comment/", controllers.AddComment)

	http.HandleFunc("/api/admin/users/update-role", IsAdmin(controllers.AdminUpdateUserRole))
}

// UTILITIES
func IsAdmin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentUser, err := controllers.GetCurrentLoggedInUser(r)
		if err != nil {
			http.Error(w, "Failed to fetch user", http.StatusUnauthorized)
			log.Printf("Error fetching user: %v", err)
			return
		}

		if currentUser.Role != "admin" {
			http.Error(w, "Not authorized to access this resource", http.StatusUnauthorized)
			return
		}

		handler(w, r)
	}
}
