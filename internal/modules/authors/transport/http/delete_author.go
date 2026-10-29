package authorshttptransport

import (
	"library/internal/core/logger"
	corehttprequest "library/internal/core/transport/http/request"
	corehttpresponse "library/internal/core/transport/http/response"
	"net/http"
)

func (h *AuthorsHTTPHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
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

	if err := h.authorsService.Delete(ctx, authorID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete author",
		)
		return
	}

	responseHandler.NoContentResponse()
}
