package controllers

import (
	"ff/database/models"
	"log"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postUUID := strings.TrimPrefix(r.URL.Path, "/api/posts/comment/")

	currentUser, err := GetCurrentLoggedInUser(r)
	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		log.Printf("Error fetching user: %v", err)
		return
	}

	author := currentUser.FirstName
	content := r.FormValue("content")

	uuid, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		log.Printf("Error generating UUID: %v", err)
		return
	}

	INSERT_QUERY := `INSERT INTO comments (UUID, PostUUID, Author, Content) VALUES (?, ?, ?, ?)`

	_, err = db.Exec(INSERT_QUERY, uuid.String(), postUUID, author, content)
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		log.Printf("Error inserting comment into database: %v", err)
		return
	}

	http.Redirect(w, r, "/posts/"+postUUID, http.StatusSeeOther)

}

func GetCommentsByPostUUID(postUUID string) ([]models.Comment, error) {
	SELECT_COMMENTS_BY_POST_UUID := "SELECT UUID, PostUUID, Author, Content, Date FROM comments WHERE PostUUID = ?"

	rows, err := db.Query(SELECT_COMMENTS_BY_POST_UUID, postUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.UUID, &comment.PostUUID, &comment.Author, &comment.Content, &comment.Date)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
