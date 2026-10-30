package authorspostgres

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
)

func (r *AuthorsRepository) Create(
	ctx context.Context,
	author authorsdomain.Author,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	const query = `
		INSERT INTO library.authors (
			id, first_name, last_name, middle_name,
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		author.ID.String(),
		author.FirstName,
		author.LastName,
		author.MiddleName,
		author.CreatedAt,
		author.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("insert author: %w", err)
	}

	return nil
}
