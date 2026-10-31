package models

import (
	booksdomain "library/internal/modules/books/domain"
	"time"
	"uuid"
)

type Book struct {
	ID          uuid.UUID
	Title       string
	Description *string
	Authors     []uuid.UUID
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b Book) ToDomain() booksdomain.Book {
	return booksdomain.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Authors:     b.Authors,
		IsAvailable: b.IsAvailable,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
