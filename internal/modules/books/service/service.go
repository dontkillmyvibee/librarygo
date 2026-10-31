package booksservice

import (
	"context"
	booksdomain "library/internal/modules/books/domain"
	"uuid"
)

type BooksService struct {
	booksRepository BooksRepository
}

type BooksRepository interface {
	Create(
		ctx context.Context,
		input CreateBookInput,
	) (booksdomain.Book, error)
	GetBook(
		ctx context.Context,
		id uuid.UUID,
	) (booksdomain.Book, error)
}

func NewBooksService(booksRepository BooksRepository) *BooksService {
	return &BooksService{
		booksRepository: booksRepository,
	}
}
