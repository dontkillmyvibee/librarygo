package authorsdomain

import (
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
