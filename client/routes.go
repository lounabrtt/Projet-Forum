package client

import (
	"ff/client/middlewares"
	"ff/client/templates"
	"net/http"
)

func ServeFileHandler(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
}

func Routes() {
	// ##### Pages #####
	http.HandleFunc("/", templates.Home)

	http.HandleFunc("/news", templates.News)
	http.HandleFunc("/news/", templates.SingleNew)
	http.HandleFunc("/news/create", middlewares.RequireRole([]string{"admin"}, templates.CreateNews))

	http.HandleFunc("/posts", templates.Posts)
	http.HandleFunc("/posts/create", middlewares.RequireRole([]string{"admin", "user"}, templates.CreatePost))
	http.HandleFunc("/posts/", templates.SinglePost)

	http.HandleFunc("/leaderboard", templates.Leaderboard)

	http.HandleFunc("/legalnotice", templates.Legalnotice)

	http.HandleFunc("/admin", ServeFileHandler("./web/pages/admin.html"))
	
	http.HandleFunc("/profile", templates.Profile)

	http.HandleFunc("/signup", ServeFileHandler("./web/pages/signup.html"))
	http.HandleFunc("/login", ServeFileHandler("./web/pages/login.html"))

	// ##### Static files #####
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("./web/pictures"))))
}
