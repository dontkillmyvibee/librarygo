package authorsservice

import (
	"context"
	"fmt"
	"uuid"
)

func (s *AuthorsService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.authorsRepository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete author from repository: %w", err)
	}

	return nil
}
