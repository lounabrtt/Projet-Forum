package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
)

type AdminTemplate struct {
	CurrentUser models.User
	LoggedIn    bool
	Users       []models.User
}

func Admin(w http.ResponseWriter, r *http.Request) {
	currentUser, err := controllers.GetCurrentLoggedInUser(r)
	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			log.Printf("Error fetching user: %v", err)
			return
		}
	}

	users, err := controllers.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		log.Printf("Error fetching users: %v", err)
		return
	}

	if currentUser == (models.User{}) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	data := AdminTemplate{
		CurrentUser: currentUser,
		LoggedIn:    currentUser.UUID != "",
		Users:       users,
	}

	tmpl, err := template.ParseFiles("web/pages/admin.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
	}
}
