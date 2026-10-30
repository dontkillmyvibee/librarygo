package authorspostgres

import (
	"context"
	"errors"
	"fmt"
	"uuid"

	coreerrors "library/internal/core/errors"
	authorsdomain "library/internal/modules/authors/domain"
	authorspostgresmodels "library/internal/modules/authors/repository/postgres/models"
	authorsservice "library/internal/modules/authors/service"

	"github.com/jackc/pgx/v5"
)

func (r *AuthorsRepository) Patch(
	ctx context.Context,
	id uuid.UUID,
	input authorsservice.PatchAuthorInput,
) (authorsdomain.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	const query = `
		UPDATE library.authors
		SET first_name = COALESCE($2, first_name),
		    last_name = COALESCE($3, last_name),
		    middle_name = CASE
		        WHEN $4 THEN $5
		        ELSE middle_name
		    END,
		    updated_at = now()
		WHERE id = $1
		RETURNING first_name, last_name, middle_name,
		          created_at, updated_at
	`

	model := authorspostgresmodels.Author{
		ID: id,
	}

	err := r.pool.QueryRow(
		ctx,
		query,
		id.String(),
		input.FirstName,
		input.LastName,
		input.MiddleNameSet,
		input.MiddleName,
	).Scan(
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
		return authorsdomain.Author{}, fmt.Errorf("patch author: %w", err)
	}

	return model.ToDomain(), nil
}
