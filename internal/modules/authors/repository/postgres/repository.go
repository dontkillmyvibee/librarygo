package authorspostgres

import (
	corepostgres "library/internal/core/repository/postgres"
)

type AuthorsRepository struct {
	pool *corepostgres.Pool
}

func NewAuthorsRepository(pool *corepostgres.Pool) *AuthorsRepository {
	return &AuthorsRepository{
		pool: pool,
	}
}
