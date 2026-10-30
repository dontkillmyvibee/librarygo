package authorspostgres

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
	authorspostgresmodels "library/internal/modules/authors/repository/postgres/models"
)

func (r *AuthorsRepository) Create(
	ctx context.Context,
	author authorsdomain.Author,
) (authorsdomain.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	const query = `
		INSERT INTO library.authors (
			first_name, last_name, middle_name
		)
		VALUES ($1, $2, $3)
		RETURNING id, first_name, last_name, middle_name, created_at, updated_at
	`

	var model authorspostgresmodels.Author

	err := r.pool.QueryRow(
		ctx,
		query,
		author.FirstName,
		author.LastName,
		author.MiddleName,
	).Scan(
		&model.ID,
		&model.FirstName,
		&model.LastName,
		&model.MiddleName,
		&model.CreatedAt,
		&model.UpdatedAt,
	)
	if err != nil {
		return authorsdomain.Author{}, fmt.Errorf(
			"insert author: %w",
			err,
		)
	}

	return model.ToDomain(), nil
}
