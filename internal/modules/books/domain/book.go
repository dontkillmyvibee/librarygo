package booksdomain

import (
	"time"
	"uuid"
)

type Book struct {
	ID          uuid.UUID
	Title       string
	Description *string
	Authors     []uuid.UUID
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
