package bookshttpdto

import (
	"library/internal/core/validation"
	"strings"
	"uuid"
)

type CreateBookRequest struct {
	Title       string      `json:"title"`
	Description *string     `json:"description"`
	Authors     []uuid.UUID `json:"authors"`
}

type CreateBookResponse struct {
	ID          uuid.UUID   `json:"id"`
	Title       string      `json:"title"`
	Description *string     `json:"description,omitempty"`
	Authors     []uuid.UUID `json:"authors"`
	IsAvailable bool        `json:"is_available"`
}

func (r *CreateBookRequest) Normalize() {
	r.Title = strings.TrimSpace(r.Title)

	if r.Description != nil {
		desc := strings.TrimSpace(*r.Description)

		if desc == "" {
			r.Description = nil
		} else {
			r.Description = &desc
		}
	}
}

func (r *CreateBookRequest) Validate() error {
	if err := validation.ValidateRequiredString(
		"title",
		r.Title,
	); err != nil {
		return err
	}

	if err := validation.ValidateStringWithLen(
		"title",
		r.Title,
		100,
	); err != nil {
		return err
	}

	if r.Description != nil {
		if err := validation.ValidateStringWithLen(
			"description",
			*r.Description,
			1000,
		); err != nil {
			return err
		}
	}

	if err := validation.ValidateUniqueUUIDsInSlice(
		"authors",
		r.Authors,
	); err != nil {
		return err
	}

	return nil
}
