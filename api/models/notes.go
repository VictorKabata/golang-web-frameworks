package models

type Note struct {
	ID          int    `json: "ID"`
	Title       string `json: "title"`
	Description string `json: "description"`
}
