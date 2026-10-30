package authorspostgres

import (
	"context"
	"fmt"
	authorsdomain "library/internal/modules/authors/domain"
	authorspostgresmodels "library/internal/modules/authors/repository/postgres/models"
)

func (r *AuthorsRepository) GetAuthors(
	ctx context.Context,
) ([]authorsdomain.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	const query = `
		select id, first_name, last_name, middle_name, created_at, updated_at
		from library.authors
		order by created_at, id
	`

	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("query authors: %w", err)
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
			return nil, fmt.Errorf("scan authors: %w", err)
		}

		authors = append(authors, model.ToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate authors: %w", err)
	}

	return authors, nil
}
