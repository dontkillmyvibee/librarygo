package authorspostgresmodels

import (
	authorsdomain "library/internal/modules/authors/domain"
	"time"
	"uuid"
)

type Author struct {
	ID         uuid.UUID
	FirstName  string
	LastName   string
	MiddleName *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (a Author) ToDomain() authorsdomain.Author {
	return authorsdomain.Author{
		ID:         a.ID,
		FirstName:  a.FirstName,
		LastName:   a.LastName,
		MiddleName: a.MiddleName,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}
