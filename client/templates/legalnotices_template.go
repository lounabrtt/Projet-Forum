package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
)

type PrivacypolicyTemplate struct {
	CurrentUser models.User
	LoggedIn    bool
}

func Privacypolicy(w http.ResponseWriter, r *http.Request) {
	currentUser, err := controllers.GetCurrentLoggedInUser(r)

	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			log.Printf("Error fetching user: %v", err)
			return
		}
	}

	data := PrivacypolicyTemplate{
		CurrentUser: currentUser,
		LoggedIn:    currentUser.UUID != "",
	}

	tmpl, err := template.ParseFiles("web/pages/privacypolicy.html")
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
