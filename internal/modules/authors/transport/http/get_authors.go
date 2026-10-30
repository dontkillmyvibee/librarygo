package authorshttptransport

import (
	"library/internal/core/logger"
	corehttprequest "library/internal/core/transport/http/request"
	corehttpresponse "library/internal/core/transport/http/response"
	authorsservice "library/internal/modules/authors/service"
	authorshttpdto "library/internal/modules/authors/transport/http/dto"
	"net/http"
)

func (h *AuthorsHTTPHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

	limit, offset, err := corehttprequest.GetLimitAndOffset(
		r,
		20,
		100,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"invalid pagination parameters",
		)
		return
	}

	authors, err := h.authorsService.GetAuthors(
		ctx,
		authorsservice.GetAuthorsInput{
			Limit:  limit,
			Offset: offset,
		},
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get authors",
		)
		return
	}

	responseHandler.JSONResponse(
		authorshttpdto.AuthorsFromService(authors),
		http.StatusOK,
	)
}
