package main

import (
	"context"
	"library/config"
	"library/internal/core/logger"
	httpserver "library/internal/core/transport/http/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log, err := logger.New(cfg.Logger)
	if err != nil {
		panic(err)
	}

	httpServer := httpserver.NewServer(cfg.HTTP, log)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("http server failed", "error", err)
		os.Exit(1)
	}

	log.Info("http server stopped")
}
