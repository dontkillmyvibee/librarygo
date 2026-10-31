package bookshttptransport

import (
	"library/internal/core/logger"
	corehttprequest "library/internal/core/transport/http/request"
	corehttpresponse "library/internal/core/transport/http/response"
	booksservice "library/internal/modules/books/service"
	bookshttpdto "library/internal/modules/books/transport/http/dto"
	"net/http"
)

func (h *BooksHTTPHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

	var req bookshttpdto.CreateBookRequest

	if err := corehttprequest.DecodeRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	req.Normalize()
	if err := req.Validate(); err != nil {
		responseHandler.ErrorResponse(
			err,
			"request validation failed",
		)
		return
	}

	book, err := h.booksService.Create(
		ctx,
		booksservice.CreateBookInput{
			Title:       req.Title,
			Description: req.Description,
			Authors:     req.Authors,
		},
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create book",
		)
		return
	}

	responseHandler.JSONResponse(
		bookshttpdto.BookFromDomain[bookshttpdto.CreateBookResponse](book),
		http.StatusCreated,
	)
}
