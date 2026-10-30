package corehttprequest

import (
	"fmt"
	"library/internal/core/validation"
	"net/http"
	"strconv"
)

func GetLimitAndOffset(
	r *http.Request,
	defaultLimit,
	maxLimit int,
) (int, int, error) {
	if defaultLimit < 1 || defaultLimit > maxLimit {
		return 0, 0, fmt.Errorf(
			"invalid pagination configuration: defaultLimit=%d, maxLimit=%d",
			defaultLimit,
			maxLimit,
		)
	}

	query := r.URL.Query()
	limit := defaultLimit
	offset := 0

	if raw := query.Get("limit"); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return 0, 0, &validation.FieldError{
				Field:   "limit",
				Message: "must be an integer",
			}
		}

		limit = value
	}

	if raw := query.Get("offset"); raw != "" {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return 0, 0, &validation.FieldError{
				Field:   "offset",
				Message: "must be an integer",
			}
		}

		offset = value
	}

	if limit < 1 || limit > maxLimit {
		return 0, 0, &validation.FieldError{
			Field:   "limit",
			Message: fmt.Sprintf("must be between 1 and %d", maxLimit),
		}
	}

	if offset < 0 {
		return 0, 0, &validation.FieldError{
			Field:   "offset",
			Message: "must be greater than or equal to 0",
		}
	}

	return limit, offset, nil
}
