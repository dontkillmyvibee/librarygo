package main

import (
	"context"
	"library/internal/core/transport/http/middleware"
	authorsmemory "library/internal/modules/authors/repository/memory"
	authorsservice "library/internal/modules/authors/service"
	authorshttptransport "library/internal/modules/authors/transport/http"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	"library/config"
	"library/internal/core/logger"
	httpserver "library/internal/core/transport/http/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		stdlog.Fatalf("load config: %v", err)
	}

	log, err := logger.New(cfg.Logger)
	if err != nil {
		stdlog.Fatalf("initialize logger: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	authorsRepository := authorsmemory.New()
	authorsService := authorsservice.NewAuthorsService(authorsRepository)
	authorsHTTPHandler := authorshttptransport.NewAuthorsHTTPHandler(authorsService)

	httpServer := httpserver.NewServer(
		cfg.HTTP,
		log,
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.AccessLog(),
		middleware.Recoverer(),
	)

	apiVersionRouterV1 := httpserver.NewApiVersionRouter(httpserver.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(authorsHTTPHandler.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run error", "error", err)
		cancel()
		os.Exit(1)
	}
}
