package corehttprequest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	coreerrors "library/internal/core/errors"
)

func DecodeRequest(r *http.Request, dest any) error {
	decoder := json.NewDecoder(r.Body)
	var raw json.RawMessage
	if err := decoder.Decode(&raw); err != nil {
		return fmt.Errorf("decode JSON body: %w: %w", coreerrors.ErrInvalidArgument, err)
	}

	if body := bytes.TrimSpace(raw); len(body) == 0 || body[0] != '{' {
		return fmt.Errorf("JSON body must be an object: %w", coreerrors.ErrInvalidArgument)
	}

	var extra json.RawMessage
	if err := decoder.Decode(&extra); err == nil {
		return fmt.Errorf("expected a single JSON object: %w", coreerrors.ErrInvalidArgument)
	} else if !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode trailing JSON data: %w: %w", coreerrors.ErrInvalidArgument, err)
	}

	fields := json.NewDecoder(bytes.NewReader(raw))
	fields.DisallowUnknownFields()
	if err := fields.Decode(dest); err != nil {
		var invalidDest *json.InvalidUnmarshalError
		if errors.As(err, &invalidDest) {
			return fmt.Errorf("invalid decode destination: %w", err)
		}

		return fmt.Errorf("decode JSON fields: %w: %w", coreerrors.ErrInvalidArgument, err)
	}

	return nil
}
