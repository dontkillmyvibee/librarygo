package authorspostgres

import (
	"context"
	"errors"
	"fmt"
	coreerrors "library/internal/core/errors"
	authorsdomain "library/internal/modules/authors/domain"
	authorspostgresmodels "library/internal/modules/authors/repository/postgres/models"
	"uuid"

	"github.com/jackc/pgx/v5"
)

func (r *AuthorsRepository) GetAuthor(
	ctx context.Context,
	id uuid.UUID,
) (authorsdomain.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	const query = `
		select first_name, last_name, middle_name, created_at, updated_at
		from library.authors
		where id = $1
	`

	model := authorspostgresmodels.Author{
		ID: id,
	}

	err := r.pool.QueryRow(ctx, query, id.String()).Scan(
		&model.FirstName,
		&model.LastName,
		&model.MiddleName,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return authorsdomain.Author{}, coreerrors.ErrNotFound
	}

	if err != nil {
		return authorsdomain.Author{}, fmt.Errorf("select author: %w", err)
	}

	return model.ToDomain(), nil
}
