package authorsservice

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
	"uuid"
)

func (s *AuthorsService) GetAuthor(
	ctx context.Context,
	id uuid.UUID,
) (authorsdomain.Author, error) {
	author, err := s.authorsRepository.GetAuthor(ctx, id)
	if err != nil {
		return authorsdomain.Author{}, fmt.Errorf(
			"get author from repository: %w",
			err,
		)
	}

	return author, nil
}
