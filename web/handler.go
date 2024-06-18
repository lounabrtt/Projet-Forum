package web

import "net/http"

func ServeFileHandler(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	}
}

func Handler() {
	// Pages 
	http.HandleFunc("/", ServeFileHandler("./index.html"))
	http.HandleFunc("/news", ServeFileHandler("./web/pages/news.html"))
	http.HandleFunc("/posts", ServeFileHandler("./web/pages/posts.html"))
	http.HandleFunc("/login", ServeFileHandler("./web/pages/login.html"))
	http.HandleFunc("/signup", ServeFileHandler("./web/pages/signup.html"))
	http.HandleFunc("/leaderboard", ServeFileHandler("./web/pages/leaderboard.html"))
	http.HandleFunc("/admin", ServeFileHandler("./web/pages/admin.html"))
	http.HandleFunc("/article", ServeFileHandler("./web/pages/news/index.html"))
	http.HandleFunc("/post", ServeFileHandler("./web/pages/posts/index.html"))
	
	
	// Static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("./web/pictures"))))
}