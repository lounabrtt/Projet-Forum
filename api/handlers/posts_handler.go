package handlers

import (
	"fmt"
	"net/http"
)


func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Create Post")
	}
}