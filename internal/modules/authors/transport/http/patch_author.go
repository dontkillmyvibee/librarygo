package authorshttptransport

import (
	"library/internal/core/logger"
	corehttprequest "library/internal/core/transport/http/request"
	corehttpresponse "library/internal/core/transport/http/response"
	authorsservice "library/internal/modules/authors/service"
	authorshttpdto "library/internal/modules/authors/transport/http/dto"
	"net/http"
)

func (h *AuthorsHTTPHandler) Patch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

	var req authorshttpdto.PatchAuthorRequest

	authorID, err := corehttprequest.GetUUIDParseValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get author ID",
		)
		return
	}

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

	author, err := h.authorsService.Patch(
		ctx,
		authorID,
		authorsservice.PatchAuthorInput{
			FirstName:     req.FirstName.Value,
			LastName:      req.LastName.Value,
			MiddleName:    req.MiddleName.Value,
			MiddleNameSet: req.MiddleName.Set,
		},
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch author",
		)
		return
	}

	responseHandler.JSONResponse(
		authorshttpdto.AuthorFromDomain[authorshttpdto.GetAuthorResponse](author),
		http.StatusOK,
	)
}
