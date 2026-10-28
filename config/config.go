package config

import (
	"fmt"
	"library/internal/core/logger"
	"library/internal/core/transport/http/server"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	HTTP   httpserver.Config
	Logger logger.Config
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
