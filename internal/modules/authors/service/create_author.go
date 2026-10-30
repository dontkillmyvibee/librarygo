package authorsservice

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
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
	createdAuthor, err := s.authorsRepository.Create(
		ctx,
		authorsdomain.Author{
			FirstName:  input.FirstName,
			LastName:   input.LastName,
			MiddleName: input.MiddleName,
		},
	)
	if err != nil {
		return authorsdomain.Author{}, fmt.Errorf("create author: %w", err)
	}

	return createdAuthor, nil
}
