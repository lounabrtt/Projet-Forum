package templates

import (
    "ff/api/controllers"
    "ff/database/models"
    "html/template"
    "log"
    "net/http"
    "strings"
)

type SinglePostTemplate struct {
    CurrentUser models.User
    LoggedIn    bool
    Post        models.Post
}

func SinglePost(w http.ResponseWriter, r *http.Request) {
    currentUser, err := controllers.GetCurrentLoggedInUser(r)
    if err != nil {
        if err != http.ErrNoCookie {
            http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
            log.Printf("Error fetching user: %v", err)
            return
        }
    }

    uuid := strings.TrimPrefix(r.URL.Path, "/posts/")
    if uuid == "" {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    post, err := controllers.GetPostByUuid(uuid)
    if err != nil {
        log.Printf("Error fetching post: %v", err)
        http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
        return
    }

    post.FormattedContent = template.HTML(strings.ReplaceAll(post.Content, "\n", "<br>"))

    data := SinglePostTemplate{
        Post:        post,
        LoggedIn:    currentUser.UUID != "",
        CurrentUser: currentUser,
    }

    tmpl, err := template.ParseFiles("web/pages/posts/index.html")
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
