package templates

import (
	"ff/api/controllers"
	"ff/database/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type SingleNewTemplate struct {
	CurrentUser models.User
	LoggedIn    bool
	New        models.New
}

func SingleNew(w http.ResponseWriter, r *http.Request) {
	currentUser, err := controllers.GetCurrentLoggedInUser(r)
	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			log.Printf("Error fetching user: %v", err)
			return
		}
	}

	uuid := strings.TrimPrefix(r.URL.Path, "/news/")
	if uuid == "" {
		http.Error(w, "Invalid new ID", http.StatusBadRequest)
		return
	}

	new, err := controllers.GetNewByUuid(uuid)
	if err != nil {
		log.Printf("Error fetching new: %v", err)
		http.Error(w, "Failed to fetch new", http.StatusInternalServerError)
		return
	}

    new.FormattedContent = template.HTML(strings.ReplaceAll(new.Content, "\n", "<br>"))

	data := SingleNewTemplate{
		New:        new,
		LoggedIn:    currentUser.UUID != "",
		CurrentUser: currentUser,
	}

	tmpl, err := template.ParseFiles("web/pages/news/index.html")
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
