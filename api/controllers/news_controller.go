package controllers

import (
	"database/sql"
	"ff/database/models"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)


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

    rows, err := db.Query("SELECT UUID, Title, Content, Date, Category, Author FROM news ORDER BY Date DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var new models.New
        err := rows.Scan(&new.UUID, &new.Title, &new.Content, &new.Date, &new.Category, &new.Author)
        if err != nil {
            return nil, err
        }
        news = append(news, new)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return news, nil
}

func GetNewByUuid(uuid string) (models.New, error) {
	SELECT_NEW_BY_UUID := "SELECT UUID, title, content, date, category, author FROM news WHERE UUID = ?"

	var new models.New
	err := db.QueryRow(SELECT_NEW_BY_UUID, uuid).Scan(&new.UUID, &new.Title, &new.Content, &new.Date, &new.Category, &new.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.New{}, nil
		}
		return models.New{}, err
	}

	return new, nil
}
