package authorsservice

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
	"uuid"
)

type PatchAuthorInput struct {
	FirstName     *string
	LastName      *string
	MiddleName    *string
	MiddleNameSet bool
}

func (s *AuthorsService) Patch(
	ctx context.Context,
	id uuid.UUID,
	input PatchAuthorInput,
) (authorsdomain.Author, error) {
	if input.FirstName == nil &&
		input.LastName == nil &&
		!input.MiddleNameSet {
		return s.GetAuthor(ctx, id)
	}

	author, err := s.authorsRepository.Patch(ctx, id, input)
	if err != nil {
		return authorsdomain.Author{}, fmt.Errorf(
			"patch author: %w",
			err,
		)
	}

	return author, nil
}
