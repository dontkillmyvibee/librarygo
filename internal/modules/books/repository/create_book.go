package booksrepository

import (
	"context"
	"errors"
	"fmt"
	coreerrors "library/internal/core/errors"
	booksdomain "library/internal/modules/books/domain"
	"library/internal/modules/books/repository/models"
	booksservice "library/internal/modules/books/service"
	"uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *BooksRepository) Create(
	ctx context.Context,
	input booksservice.CreateBookInput,
) (booksdomain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return booksdomain.Book{}, fmt.Errorf("begin create book transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	const queryInsertBook = `
		insert into library.books (title, description)
		values ($1, $2)
		returning id, title, description, is_available, created_at, updated_at
	`

	var modelBook models.Book

	if err := tx.QueryRow(
		ctx,
		queryInsertBook,
		input.Title,
		input.Description,
	).Scan(
		&modelBook.ID,
		&modelBook.Title,
		&modelBook.Description,
		&modelBook.IsAvailable,
		&modelBook.CreatedAt,
		&modelBook.UpdatedAt,
	); err != nil {
		return booksdomain.Book{}, fmt.Errorf(
			"insert book: %w",
			err,
		)
	}

	const queryInsertAuthors = `
    	insert into library.book_authors (book_id, author_id)
    	values ($1, $2)
    `

	for _, authorID := range input.Authors {
		_, err := tx.Exec(
			ctx,
			queryInsertAuthors,
			modelBook.ID,
			authorID,
		)
		if err != nil {
			var pgErr *pgconn.PgError

			if errors.As(err, &pgErr) &&
				pgErr.Code == "23503" &&
				pgErr.ConstraintName == "book_authors_author_id_fkey" {
				return booksdomain.Book{}, fmt.Errorf(
					"author %s does not exist: %w: %w",
					authorID.String(),
					coreerrors.ErrInvalidArgument,
					err,
				)
			}

			return booksdomain.Book{}, fmt.Errorf(
				"insert book-author: %w",
				err,
			)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return booksdomain.Book{}, fmt.Errorf(
			"commit create book transaction: %w",
			err,
		)
	}

	modelBook.Authors = append([]uuid.UUID(nil), input.Authors...)

	return modelBook.ToDomain(), nil
}
