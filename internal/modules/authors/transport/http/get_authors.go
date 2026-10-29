package authorshttptransport

import (
	"library/internal/core/logger"
	corehttpresponse "library/internal/core/transport/http/response"
	authorshttpdto "library/internal/modules/authors/transport/http/dto"
	"net/http"
)

func (h *AuthorsHTTPHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

	authors, err := h.authorsService.GetAuthors(ctx)
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
