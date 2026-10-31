package booksservice

import (
	"context"
	"fmt"
	"library/internal/core/validation"
	booksdomain "library/internal/modules/books/domain"
	"uuid"
)

type CreateBookInput struct {
	Title       string
	Description *string
	Authors     []uuid.UUID
}

func (s *BooksService) Create(
	ctx context.Context,
	input CreateBookInput,
) (booksdomain.Book, error) {
	if len(input.Authors) == 0 {
		return booksdomain.Book{}, &validation.FieldError{
			Field:   "authors",
			Message: "authors must not be empty",
		}
	}

	book, err := s.booksRepository.Create(ctx, input)
	if err != nil {
		return booksdomain.Book{}, fmt.Errorf("create book: %w", err)
	}

	return book, nil
}
