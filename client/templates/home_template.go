package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	currentUser, err := controllers.GetCurrentLoggedInUser(r)
	if err != nil {
		if err == http.ErrNoCookie {
			data := struct {
				CurrentUser models.User
				Connected   bool
			}{
				CurrentUser: models.User{},
				Connected:   false,
			}

			tmpl, err := template.ParseFiles("index.html")
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
			return
		}

		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		log.Printf("Error fetching user: %v", err)
		return
	}

	data := struct {
		CurrentUser models.User
		Connected   bool
	}{
		CurrentUser: currentUser,
		Connected:   true,
	}

	tmpl, err := template.ParseFiles("index.html")
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
