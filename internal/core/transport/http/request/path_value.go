package corehttprequest

import (
	"fmt"
	coreerrors "library/internal/core/errors"
	"net/http"
	"uuid"
)

func GetUUIDParseValue(r *http.Request, key string) (uuid.UUID, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return uuid.Nil(), fmt.Errorf(
			"get path value by key %s is empty: %w",
			key,
			coreerrors.ErrInvalidArgument,
		)
	}

	val, err := uuid.Parse(pathValue)
	if err != nil {
		return uuid.Nil(), fmt.Errorf(
			"parse path parameter %q as UUID: %w: %w",
			key,
			coreerrors.ErrInvalidArgument,
			err,
		)
	}

	return val, nil
}
