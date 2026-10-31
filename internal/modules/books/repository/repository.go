package booksrepository

import corepostgres "library/internal/core/repository/postgres"

type BooksRepository struct {
	pool *corepostgres.Pool
}

func NewBooksRepository(pool *corepostgres.Pool) *BooksRepository {
	return &BooksRepository{
		pool: pool,
	}
}
