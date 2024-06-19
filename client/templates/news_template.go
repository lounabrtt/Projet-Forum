package templates

import (
    "ff/api/controllers"
    "ff/database/models"
    "html/template"
    "log"
    "net/http"
    "strings"
)

type NewsTemplate struct {
    CurrentUser models.User
    LoggedIn    bool
    News        []models.New
}

// Word limit on News content 
func limitWords(s string, limit int) string {
    words := strings.Fields(s)
    if len(words) > limit {
        return strings.Join(words[:limit], " ") + "... see more" 
    }
    return s
}

func News(w http.ResponseWriter, r *http.Request) {
    currentUser, err := controllers.GetCurrentLoggedInUser(r)
    if err != nil {
        if err != http.ErrNoCookie {
            http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
            log.Printf("Error fetching user: %v", err)
            return
        }
    }

    news, err := controllers.GetAllNews()
    if err != nil {
        http.Error(w, "Failed to fetch news", http.StatusInternalServerError)
        log.Printf("Error fetching news: %v", err)
        return
    }

    for i := range news {
        news[i].Content = limitWords(news[i].Content, 50) // limit to 50 words 
    }

    data := NewsTemplate{
        CurrentUser: currentUser,
        LoggedIn:    currentUser.UUID != "",
        News:        news,
    }

    tmpl, err := template.ParseFiles("web/pages/news.html")
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
