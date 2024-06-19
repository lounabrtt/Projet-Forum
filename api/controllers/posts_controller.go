package controllers

import (
	"database/sql"
	"ff/database/models"
	"log"
	"net/http"
	"sort"

	"github.com/gofrs/uuid"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content") 
	author := r.FormValue("author")
	category := r.FormValue("category")

	uuid, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		log.Printf("Error generating UUID: %v", err)
		return
	}

	INSERT_QUERY := `INSERT INTO posts (UUID, title, content, author, category) VALUES (?, ?, ?, ?, ?)`

	_, err = db.Exec(INSERT_QUERY, uuid.String(), title, content, author, category)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		log.Printf("Error inserting post into database: %v", err)
		return
	}

	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}

func GetAllPosts() (models.Posts, error) {
	var posts models.Posts

	rows, err := db.Query("SELECT UUID, Title, Content, Author, Category, Date FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.UUID, &post.Title, &post.Content, &post.Author, &post.Category, &post.Date)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPostByUuid(uuid string) (models.Post, error) {
	SELECT_POST_BY_UUID := "SELECT UUID, title, content, author FROM posts WHERE UUID = ?"

	var post models.Post
	err := db.QueryRow(SELECT_POST_BY_UUID, uuid).Scan(&post.UUID, &post.Title, &post.Content, &post.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, nil
		}
		return models.Post{}, err
	}

	comments, err := GetCommentsByPostUUID(post.UUID)
	if err != nil {
		return models.Post{}, err
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Date > comments[j].Date
	})

	post.Comments = comments

	return post, nil
}
