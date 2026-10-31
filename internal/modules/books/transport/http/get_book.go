package bookshttptransport

import (
	"library/internal/core/logger"
	corehttprequest "library/internal/core/transport/http/request"
	corehttpresponse "library/internal/core/transport/http/response"
	bookshttpdto "library/internal/modules/books/transport/http/dto"
	"net/http"
)

func (h *BooksHTTPHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

	bookID, err := corehttprequest.GetUUIDParseValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get bookID path value",
		)
		return
	}

	book, err := h.booksService.GetBook(ctx, bookID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get book",
		)
		return
	}

	responseHandler.JSONResponse(
		bookshttpdto.BookFromDomain[bookshttpdto.GetBookResponse](book),
		http.StatusOK,
	)
}
