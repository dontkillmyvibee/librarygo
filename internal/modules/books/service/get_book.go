package booksservice

import (
	"context"
	"fmt"
	booksdomain "library/internal/modules/books/domain"
	"uuid"
)

func (s *BooksService) GetBook(
	ctx context.Context,
	id uuid.UUID,
) (booksdomain.Book, error) {
	book, err := s.booksRepository.GetBook(ctx, id)
	if err != nil {
		return booksdomain.Book{}, fmt.Errorf(
			"get book: %w",
			err,
		)
	}

	return book, nil
}
