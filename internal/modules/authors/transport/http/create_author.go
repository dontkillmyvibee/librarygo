package authorshttptransport

import (
	"library/internal/core/logger"
	corehttprequest "library/internal/core/transport/http/request"
	corehttpresponse "library/internal/core/transport/http/response"
	authorsservice "library/internal/modules/authors/service"
	authorshttpdto "library/internal/modules/authors/transport/http/dto"
	"net/http"
)

func (h *AuthorsHTTPHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

	var req authorshttpdto.CreateAuthorRequest

	if err := corehttprequest.DecodeRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	req.Normalize()
	if err := req.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "request validation failed")
		return
	}

	author, err := h.authorsService.Create(
		ctx,
		authorsservice.CreateAuthorInput{
			FirstName:  req.FirstName,
			LastName:   req.LastName,
			MiddleName: req.MiddleName,
		},
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create author")
		return
	}

	responseHandler.JSONResponse(
		authorshttpdto.AuthorFromDomain[authorshttpdto.CreateAuthorResponse](author),
		http.StatusCreated,
	)
}
