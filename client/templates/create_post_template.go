package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
)

type CreatePostTemplate struct {
	CurrentUser models.User
	LoggedIn    bool
	Categories  []models.Category 
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	currentUser, err := controllers.GetCurrentLoggedInUser(r)
	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			log.Printf("Error fetching user: %v", err)
			return
		}
	}

	categories, err := controllers.GetCategories()
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		log.Printf("Error fetching categories: %v", err)
		return
	}

	data := CreatePostTemplate{
		CurrentUser: currentUser,
		LoggedIn:    currentUser.UUID != "",
		Categories:  categories,
	}

	tmpl, err := template.ParseFiles("web/pages/posts/create.html")
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
