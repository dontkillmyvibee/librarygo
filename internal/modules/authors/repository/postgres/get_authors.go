package authorspostgres

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
	authorspostgresmodels "library/internal/modules/authors/repository/postgres/models"
	authorsservice "library/internal/modules/authors/service"

	"github.com/jackc/pgx/v5"
)

func (r *AuthorsRepository) GetAuthors(
	ctx context.Context,
	input authorsservice.GetAuthorsInput,
) ([]authorsdomain.Author, int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	tx, err := r.pool.BeginTx(
		ctx,
		pgx.TxOptions{
			IsoLevel:   pgx.RepeatableRead,
			AccessMode: pgx.ReadOnly,
		},
	)
	if err != nil {
		defer tx.Rollback(ctx)
		return nil, 0, fmt.Errorf("begin authors transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	var total int

	if err := tx.QueryRow(
		ctx,
		`select count(*) from library.authors`,
	).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("select authors count: %w", err)
	}

	const query = `
		select id, first_name, last_name, middle_name, created_at, updated_at
		from library.authors
		order by created_at, id
		limit $1 offset $2
	`

	rows, err := tx.Query(ctx, query, input.Limit, input.Offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query authors: %w", err)
	}
	defer rows.Close()

	authors := make([]authorsdomain.Author, 0)

	for rows.Next() {
		var model authorspostgresmodels.Author

		if err := rows.Scan(
			&model.ID,
			&model.FirstName,
			&model.LastName,
			&model.MiddleName,
			&model.CreatedAt,
			&model.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("scan author: %w", err)
		}

		authors = append(authors, model.ToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("iterate authors: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, 0, fmt.Errorf("commit authors tx: %w", err)
	}

	return authors, total, nil
}
