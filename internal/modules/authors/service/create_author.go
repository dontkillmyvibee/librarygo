package authorsservice

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
	"time"
	"uuid"
)

type CreateAuthorInput struct {
	FirstName  string
	LastName   string
	MiddleName *string
}

func (s *AuthorsService) Create(
	ctx context.Context,
	input CreateAuthorInput,
) (authorsdomain.Author, error) {
	now := time.Now().UTC()

	author := authorsdomain.Author{
		ID:         uuid.NewV4(),
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		MiddleName: input.MiddleName,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := s.authorsRepository.Create(ctx, author); err != nil {
		return authorsdomain.Author{}, fmt.Errorf("create author: %w", err)
	}

	return author, nil
}
