package config

import (
	"fmt"
	"library/internal/core/logger"
	corepostgres "library/internal/core/repository/postgres"
	httpserver "library/internal/core/transport/http/server"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTP     httpserver.Config
	Logger   logger.Config
	Postgres corepostgres.Config
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
