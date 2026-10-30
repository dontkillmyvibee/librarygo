package authorsservice

import (
	"context"
	authorsdomain "library/internal/modules/authors/domain"
	"time"
	"uuid"
)

type AuthorsService struct {
	authorsRepository AuthorsRepository
}

type AuthorsRepository interface {
	Create(
		ctx context.Context,
		author authorsdomain.Author,
	) error
	GetAuthor(
		ctx context.Context,
		id uuid.UUID,
	) (authorsdomain.Author, error)
	GetAuthors(
		ctx context.Context,
		input GetAuthorsInput,
	) ([]authorsdomain.Author, int, error)
	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
	Patch(
		ctx context.Context,
		id uuid.UUID,
		input PatchAuthorInput,
		updatedAt time.Time,
	) (authorsdomain.Author, error)
}

func NewAuthorsService(authorsRepository AuthorsRepository) *AuthorsService {
	return &AuthorsService{
		authorsRepository: authorsRepository,
	}
}
