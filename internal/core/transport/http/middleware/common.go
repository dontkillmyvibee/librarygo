package middleware

import (
	"library/internal/core/logger"
	corehttpresponse "library/internal/core/transport/http/response"
	"net/http"
	"time"
	"uuid"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.NewV4().String()

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := log.With(
				"request_id", r.Header.Get(requestIDHeader),
				"method", r.Method,
				"endpoint", r.URL.String(),
			)

			ctx := logger.ToContext(r.Context(), l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AccessLog() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)

			rw := corehttpresponse.NewResponseWriter(w)

			before := time.Now()
			log.Info(
				">>> incoming HTTP request",
			)

			next.ServeHTTP(rw, r)

			log.Info(
				"<<< done HTTP request",
				"status_code", rw.StatusCode(),
				"duration_ms", float64(time.Since(before))/float64(time.Millisecond),
			)
		})
	}
}

func Recoverer() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			responseHandler := corehttpresponse.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "internal server error")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
