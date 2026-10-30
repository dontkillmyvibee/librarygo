package authorshttptransport

import (
	"context"
	httpserver "library/internal/core/transport/http/server"
	authorsdomain "library/internal/modules/authors/domain"
	authorsservice "library/internal/modules/authors/service"
	"uuid"
)

type AuthorsHTTPHandler struct {
	authorsService AuthorsService
}

type AuthorsService interface {
	Create(
		ctx context.Context,
		input authorsservice.CreateAuthorInput,
	) (authorsdomain.Author, error)
	GetAuthor(
		ctx context.Context,
		id uuid.UUID,
	) (authorsdomain.Author, error)
	GetAuthors(
		ctx context.Context,
		input authorsservice.GetAuthorsInput,
	) (authorsservice.GetAuthorsOutput, error)
	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
	Patch(
		ctx context.Context,
		id uuid.UUID,
		input authorsservice.PatchAuthorInput,
	) (authorsdomain.Author, error)
}

func NewAuthorsHTTPHandler(authorsService AuthorsService) *AuthorsHTTPHandler {
	return &AuthorsHTTPHandler{
		authorsService: authorsService,
	}
}

func (h *AuthorsHTTPHandler) Routes() []httpserver.Route {
	return []httpserver.Route{
		{
			Method:  httpserver.MethodGet,
			Path:    "/authors/{id}",
			Handler: h.GetAuthor,
		},
		{
			Method:  httpserver.MethodGet,
			Path:    "/authors",
			Handler: h.GetAuthors,
		},
		{
			Method:  httpserver.MethodPost,
			Path:    "/authors",
			Handler: h.CreateAuthor,
		},
		{
			Method:  httpserver.MethodDelete,
			Path:    "/authors/{id}",
			Handler: h.DeleteAuthor,
		},
		{
			Method:  httpserver.MethodPatch,
			Path:    "/authors/{id}",
			Handler: h.Patch,
		},
	}
}
