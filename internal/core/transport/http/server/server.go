package corehttpserver

import (
	"context"
	"errors"
	"fmt"
	"library/internal/core/logger"
	"library/internal/core/transport/http/middleware"
	"net/http"
)

type Server struct {
	mux        *http.ServeMux
	config     Config
	logger     *logger.Logger
	middleware []middleware.Middleware
}

func NewServer(cfg Config, log *logger.Logger, middleware ...middleware.Middleware) *Server {
	return &Server{
		mux:        http.NewServeMux(),
		config:     cfg,
		logger:     log,
		middleware: middleware,
	}
}

func (s *Server) RegisterAPIRouters(routers ...*ApiVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		s.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}

func (s *Server) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		path := string(route.Method) + " " + route.Path
		handler := route.WithMiddleware()

		s.mux.Handle(path, handler)
	}
}

func (s *Server) Run(ctx context.Context) error {
	mux := middleware.ChainMiddleware(s.mux, s.middleware...)

	server := &http.Server{
		Addr:              ":" + s.config.Port,
		Handler:           mux,
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
