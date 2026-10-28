package httpserver

import (
	"context"
	"errors"
	"fmt"
	"library/internal/core/logger"
	"net/http"
)

type Server struct {
	mux    *http.ServeMux
	config Config
	logger *logger.Logger
}

func NewServer(cfg Config, log *logger.Logger) *Server {
	return &Server{
		mux:    http.NewServeMux(),
		config: cfg,
		logger: log,
	}
}

func (s *Server) RegisterAPIRouters(routers ...ApiVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		s.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}

func (s *Server) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:              ":" + s.config.Port,
		Handler:           s.mux,
		ReadTimeout:       s.config.ReadTimeout,
		ReadHeaderTimeout: s.config.ReadHeaderTimeout,
		WriteTimeout:      s.config.WriteTimeout,
		IdleTimeout:       s.config.IdleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		defer close(serverErrors)

		s.logger.Info("starting http server", "port", s.config.Port)

		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		if err != nil {
			return fmt.Errorf("listen and serve HTTP: %w", err)
		}
	case <-ctx.Done():
		s.logger.Info("shutting down http server")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			s.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}
	}

	return nil
}
