package bookshttptransport

import (
	"context"
	httpserver "library/internal/core/transport/http/server"
	booksdomain "library/internal/modules/books/domain"
	booksservice "library/internal/modules/books/service"
)

type BooksHTTPHandler struct {
	booksService BooksService
}

type BooksService interface {
	Create(
		ctx context.Context,
		input booksservice.CreateBookInput,
	) (booksdomain.Book, error)
}

func NewBooksHTTPHandler(booksService BooksService) *BooksHTTPHandler {
	return &BooksHTTPHandler{
		booksService: booksService,
	}
}

func (h *BooksHTTPHandler) Routes() []httpserver.Route {
	return []httpserver.Route{
		{
			Method:  httpserver.MethodPost,
			Path:    "/books",
			Handler: h.CreateBook,
		},
	}
}
