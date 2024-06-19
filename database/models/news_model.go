package models

import "html/template"


type New struct {
	UUID     string `json:"uuid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Date     string `json:"date"`
	FormattedContent template.HTML `json:"-"`
}

type News []New
