package authorsservice

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
)

type GetAuthorsInput struct {
	Limit  int
	Offset int
}
type GetAuthorsOutput struct {
	Authors []authorsdomain.Author
	Total   int
}

func (s *AuthorsService) GetAuthors(
	ctx context.Context,
	input GetAuthorsInput,
) (GetAuthorsOutput, error) {
	authors, total, err := s.authorsRepository.GetAuthors(ctx, input)
	if err != nil {
		return GetAuthorsOutput{}, fmt.Errorf(
			"get authors from repository: %w",
			err,
		)
	}

	return GetAuthorsOutput{
		Authors: authors,
		Total:   total,
	}, nil
}
