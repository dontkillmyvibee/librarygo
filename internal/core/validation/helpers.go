package validation

import (
	"fmt"
	"library/internal/core/transport/http/types"
	"strings"
	"unicode"
	"unicode/utf8"
	"uuid"
)

func ValidateStringWithLen(
	fieldName,
	value string,
	maxFieldLength int,
) error {
	if utf8.RuneCountInString(value) > maxFieldLength {
		return &FieldError{
			Field:   fieldName,
			Message: fmt.Sprintf("must not exceed %d characters", maxFieldLength),
		}
	}

	if strings.ContainsFunc(value, unicode.IsControl) {
		return &FieldError{
			Field:   fieldName,
			Message: "must not contain control characters",
		}
	}

	return nil
}

func ValidateRequiredNil(fieldName string, value any) error {
	if value == nil {
		return &FieldError{
			Field:   fieldName,
			Message: "value is required",
		}
	}

	return nil
}

func ValidateRequiredString(fieldName, value string) error {
	if value == "" {
		return &FieldError{
			Field:   fieldName,
			Message: "field is required",
		}
	}

	return nil
}

func ValidateUniqueUUIDsInSlice(fieldName string, value []uuid.UUID) error {
	if len(value) == 0 {
		return &FieldError{
			Field:   fieldName,
			Message: "must contains min 1 ID",
		}
	}

	set := make(map[uuid.UUID]struct{}, len(value))

	for _, value := range value {
		if _, ok := set[value]; ok {
			return &FieldError{
				Field:   fieldName,
				Message: "all values must be unique",
			}
		}

		if value == uuid.Nil() {
			return &FieldError{
				Field:   fieldName,
				Message: "uuid can`t be nil",
			}
		}

		set[value] = struct{}{}
	}

	return nil
}

func NormalizePatchString(field *types.Nullable[string]) {
	if !field.Set || field.Value == nil {
		return
	}

	value := strings.TrimSpace(*field.Value)

	if value == "" {
		field.Set = false
		field.Value = nil
		return
	}

	field.Value = &value
}
