package controllers

import (
	"database/sql"
	"ff/database/models"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func formatDate(input string) (string, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", err
	}
	return t.Format("2006-01-02"), nil
}

func CreateNew(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	date := r.FormValue("date")
	category := r.FormValue("category")
	author := r.FormValue("author")

	uuid, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Failed to create new", http.StatusInternalServerError)
		log.Printf("Error generating UUID: %v", err)
		return
	}

	INSERT_QUERY := `INSERT INTO news (UUID, title, content, date, category, author) VALUES (?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(INSERT_QUERY, uuid.String(), title, content, date, category, author)
	if err != nil {
		http.Error(w, "Failed to create new", http.StatusInternalServerError)
		log.Printf("Error inserting new into database: %v", err)
		return
	}

	http.Redirect(w, r, "/news", http.StatusSeeOther)
}

func GetAllNews() ([]models.New, error) {
	var news []models.New

	rows, err := db.Query("SELECT UUID, Title, Content, Date, Author FROM news ORDER BY Date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var new models.New
		err := rows.Scan(&new.UUID, &new.Title, &new.Content, &new.Date, &new.Author)
		if err != nil {
			return nil, err
		}

		formattedDate, err := formatDate(new.Date)
		if err != nil {
			return nil, err
		}
		new.Date = formattedDate

		news = append(news, new)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}

func GetNewByUuid(uuid string) (models.New, error) {
	SELECT_NEW_BY_UUID := "SELECT UUID, title, content, date, author FROM news WHERE UUID = ?"

	var new models.New
	err := db.QueryRow(SELECT_NEW_BY_UUID, uuid).Scan(&new.UUID, &new.Title, &new.Content, &new.Date, &new.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.New{}, nil
		}
		return models.New{}, err
	}

	formattedDate, err := formatDate(new.Date)
	if err != nil {
		return models.New{}, err
	}
	new.Date = formattedDate

	return new, nil
}
