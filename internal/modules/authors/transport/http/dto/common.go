package authorshttpdto

import (
	"library/internal/core/transport/http/types"
	"library/internal/core/validation"
	authorsdomain "library/internal/modules/authors/domain"
	authorsservice "library/internal/modules/authors/service"
	"strings"
	"uuid"
)

type CreateAuthorRequest struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
}

type CreateAuthorResponse struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName *string   `json:"middle_name,omitempty"`
}

type GetAuthorResponse CreateAuthorResponse
type GetAuthorsResponse struct {
	Authors []GetAuthorResponse `json:"authors"`
	Total   int                 `json:"total"`
}

type PatchAuthorRequest struct {
	FirstName  types.Nullable[string] `json:"first_name"`
	LastName   types.Nullable[string] `json:"last_name"`
	MiddleName types.Nullable[string] `json:"middle_name"`
}

func (p *PatchAuthorRequest) Normalize() {
	validation.NormalizePatchString(&p.FirstName)
	validation.NormalizePatchString(&p.LastName)
	validation.NormalizePatchString(&p.MiddleName)
}

func (p *PatchAuthorRequest) Validate() error {
	if p.FirstName.Set {
		if p.FirstName.Value == nil {
			return &validation.FieldError{
				Field:   "first_name",
				Message: "must not be null",
			}
		}

		if err := validation.ValidateName(
			"first_name",
			*p.FirstName.Value,
			100,
		); err != nil {
			return err
		}
	}

	if p.LastName.Set {
		if p.LastName.Value == nil {
			return &validation.FieldError{
				Field:   "last_name",
				Message: "must not be null",
			}
		}

		if err := validation.ValidateName(
			"last_name",
			*p.LastName.Value,
			100,
		); err != nil {
			return err
		}
	}

	if p.MiddleName.Set && p.MiddleName.Value != nil {
		return validation.ValidateName(
			"middle_name",
			*p.MiddleName.Value,
			100,
		)
	}

	return nil
}

func AuthorFromDomain[T GetAuthorResponse | CreateAuthorResponse](author authorsdomain.Author) *T {
	return &T{
		ID:         author.ID,
		FirstName:  author.FirstName,
		LastName:   author.LastName,
		MiddleName: author.MiddleName,
	}
}

func AuthorsFromService(result authorsservice.GetAuthorsOutput) GetAuthorsResponse {
	items := make([]GetAuthorResponse, 0, len(result.Authors))

	for _, author := range result.Authors {
		items = append(items,
			*AuthorFromDomain[GetAuthorResponse](author),
		)
	}

	return GetAuthorsResponse{
		Authors: items,
		Total:   result.Total,
	}
}

func (r *CreateAuthorRequest) Normalize() {
	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)

	if r.MiddleName != nil {
		middleName := strings.TrimSpace(*r.MiddleName)

		if middleName == "" {
			r.MiddleName = nil
		} else {
			r.MiddleName = &middleName
		}
	}
}

func (r *CreateAuthorRequest) Validate() error {
	if err := validation.ValidateRequiredString("first_name", r.FirstName); err != nil {
		return err
	}

	if err := validation.ValidateName(
		"first_name",
		r.FirstName,
		100,
	); err != nil {
		return err
	}

	if err := validation.ValidateRequiredString("last_name", r.LastName); err != nil {
		return err
	}

	if err := validation.ValidateName(
		"last_name",
		r.LastName,
		100,
	); err != nil {
		return err
	}

	if r.MiddleName != nil {
		if err := validation.ValidateName(
			"middle_name",
			*r.MiddleName,
			100,
		); err != nil {
			return err
		}

	}

	return nil
}
