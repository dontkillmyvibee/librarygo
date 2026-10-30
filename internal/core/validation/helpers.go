package validation

import (
	"fmt"
	"library/internal/core/transport/http/types"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ValidateName(fieldName, value string, maxNameLength int) error {
	if utf8.RuneCountInString(value) > maxNameLength {
		return &FieldError{
			Field:   fieldName,
			Message: fmt.Sprintf("must not exceed %d characters", maxNameLength),
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
