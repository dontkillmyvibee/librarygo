package coreerrors

import "errors"

var (
	ErrForbidden            = errors.New("forbidden")
	ErrTooManyRequests      = errors.New("too many requests")
	ErrUnauthenticated      = errors.New("unauthenticated")
	ErrUnsupportedMediaType = errors.New("unsupported media type")
	ErrNotFound             = errors.New("not found")
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrConflict             = errors.New("conflict")
	ErrInternalServerError  = errors.New("internal server error")
)
