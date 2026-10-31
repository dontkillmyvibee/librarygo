package booksrepository

import (
	"context"
	"errors"
	"fmt"
	coreerrors "library/internal/core/errors"
	booksdomain "library/internal/modules/books/domain"
	"library/internal/modules/books/repository/models"
	"uuid"

	"github.com/jackc/pgx/v5"
)

func (r *BooksRepository) GetBook(
	ctx context.Context,
	id uuid.UUID,
) (booksdomain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout)
	defer cancel()

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadOnly,
	})
	if err != nil {
		return booksdomain.Book{}, fmt.Errorf(
			"begin books transaction: %w",
			err,
		)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	const queryGetBook = `
		select id, title, description, is_available, created_at, updated_at
		from library.books
		where id = $1
	`

	var modelBook models.Book

	if err := tx.QueryRow(ctx, queryGetBook, id).Scan(
		&modelBook.ID,
		&modelBook.Title,
		&modelBook.Description,
		&modelBook.IsAvailable,
		&modelBook.CreatedAt,
		&modelBook.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return booksdomain.Book{}, coreerrors.ErrNotFound
		}
		return booksdomain.Book{}, fmt.Errorf(
			"select book: %w",
			err,
		)
	}

	authorIDs := make([]uuid.UUID, 0)

	const queryBookAuthors = `
	select author_id
	from library.book_authors
	where book_id = $1
	order by author_id
	`

	rows, err := tx.Query(ctx, queryBookAuthors, id)
	if err != nil {
		return booksdomain.Book{}, fmt.Errorf(
			"select authors: %w",
			err,
		)
	}
	defer rows.Close()

	for rows.Next() {
		var authorID uuid.UUID

		if err := rows.Scan(&authorID); err != nil {
			return booksdomain.Book{}, fmt.Errorf(
				"scan book author: %w",
				err,
			)
		}

		authorIDs = append(authorIDs, authorID)
	}

	if err := rows.Err(); err != nil {
		return booksdomain.Book{}, fmt.Errorf(
			"iterate authors: %w",
			err,
		)
	}

	if err := tx.Commit(ctx); err != nil {
		return booksdomain.Book{}, fmt.Errorf(
			"commit tx: %w",
			err,
		)
	}

	modelBook.Authors = authorIDs

	return modelBook.ToDomain(), nil
}
