package authorshttptransport

import (
	"library/internal/core/logger"
	corehttprequest "library/internal/core/transport/http/request"
	corehttpresponse "library/internal/core/transport/http/response"
	authorshttpdto "library/internal/modules/authors/transport/http/dto"
	"net/http"
)

func (h *AuthorsHTTPHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

	authorID, err := corehttprequest.GetUUIDParseValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get authorID path value",
		)
		return
	}

	author, err := h.authorsService.GetAuthor(ctx, authorID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get author",
		)
		return
	}

	responseHandler.JSONResponse(
		authorshttpdto.AuthorFromDomain[authorshttpdto.GetAuthorResponse](author),
		http.StatusOK,
	)
}
