package bookshttpdto

import (
	booksdomain "library/internal/modules/books/domain"
)

func BookFromDomain[T GetBookResponse | CreateBookResponse](bookDomain booksdomain.Book) T {
	return T{
		ID:          bookDomain.ID,
		Title:       bookDomain.Title,
		Description: bookDomain.Description,
		Authors:     bookDomain.Authors,
		IsAvailable: bookDomain.IsAvailable,
	}
}
