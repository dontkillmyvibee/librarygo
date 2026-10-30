package corehttpresponse

import (
	"encoding/json"
	"errors"
	"fmt"
	coreerrors "library/internal/core/errors"
	"library/internal/core/logger"
	"library/internal/core/validation"
	"net/http"
	"runtime/debug"
)

type HTTPResponseHandler struct {
	log *logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(log *logger.Logger, w http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		w:   w,
	}
}

func (h *HTTPResponseHandler) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response", "error", err)
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		fieldErr   *validation.FieldError
		statusCode int
		publicErr  error
		logFunc    func(string, ...any)
	)
	switch {
	case errors.As(err, &fieldErr):
		statusCode, publicErr, logFunc = http.StatusBadRequest, fieldErr, h.log.Debug
	case errors.Is(err, coreerrors.ErrForbidden):
		statusCode, publicErr, logFunc = http.StatusForbidden, coreerrors.ErrForbidden, h.log.Debug
	case errors.Is(err, coreerrors.ErrTooManyRequests):
		statusCode, publicErr, logFunc = http.StatusTooManyRequests, coreerrors.ErrTooManyRequests, h.log.Debug
	case errors.Is(err, coreerrors.ErrUnauthenticated):
		statusCode, publicErr, logFunc = http.StatusUnauthorized, coreerrors.ErrUnauthenticated, h.log.Debug
	case errors.Is(err, coreerrors.ErrUnsupportedMediaType):
		statusCode, publicErr, logFunc = http.StatusUnsupportedMediaType, coreerrors.ErrUnsupportedMediaType, h.log.Debug
	case errors.Is(err, coreerrors.ErrInvalidArgument):
		statusCode, publicErr, logFunc = http.StatusBadRequest, coreerrors.ErrInvalidArgument, h.log.Debug
	case errors.Is(err, coreerrors.ErrNotFound):
		statusCode, publicErr, logFunc = http.StatusNotFound, coreerrors.ErrNotFound, h.log.Debug
	case errors.Is(err, coreerrors.ErrConflict):
		statusCode, publicErr, logFunc = http.StatusConflict, coreerrors.ErrConflict, h.log.Debug
	default:
		statusCode, publicErr, logFunc = http.StatusInternalServerError, coreerrors.ErrInternalServerError, h.log.Error
	}
	logFunc(msg, "error", err)
	h.errorResponse(statusCode, publicErr, msg)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	h.log.Error(msg,
		"panic", p,
		"panic_type", fmt.Sprintf("%T", p),
		"stack", string(debug.Stack()),
	)

	h.errorResponse(
		http.StatusInternalServerError,
		coreerrors.ErrInternalServerError,
		msg,
	)
}

func (h *HTTPResponseHandler) errorResponse(statusCode int, err error, msg string) {
	h.JSONResponse(ErrorResponse{Error: err.Error(), Message: msg}, statusCode)
}
