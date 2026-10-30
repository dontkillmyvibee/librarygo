package authorspostgres

import (
	"context"
	"fmt"
	coreerrors "library/internal/core/errors"
	"uuid"
)

func (r *AuthorsRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	const query = `
		DELETE FROM library.authors
		WHERE id = $1
	`

	tag, err := r.pool.Exec(ctx, query, id.String())
	if err != nil {
		return fmt.Errorf("delete author: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return coreerrors.ErrNotFound
	}

	return nil
}
