package entities

import "github.com/google/uuid"

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title`
	Category string `json:"category"`
}

func NewBook() *Book {
	book := Book{
		ID: uuid.New().String(),
	}

	return &book
}
