package authorsservice

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
)

type GetAuthorsOutput struct {
	Authors []authorsdomain.Author
	Total   int
}

func (s *AuthorsService) GetAuthors(ctx context.Context) (GetAuthorsOutput, error) {
	authors, err := s.authorsRepository.GetAuthors(ctx)
	if err != nil {
		return GetAuthorsOutput{}, fmt.Errorf(
			"get authors from repository: %w",
			err,
		)
	}

	total := len(authors)

	return GetAuthorsOutput{
		Authors: authors,
		Total:   total,
	}, nil
}
