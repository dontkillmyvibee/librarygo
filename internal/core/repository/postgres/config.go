package corepostgres

import (
	"net"
	"net/url"
	"time"
)

type Config struct {
	Host     string        `envconfig:"POSTGRES_HOST"     required:"true"`
	Port     string        `envconfig:"POSTGRES_PORT"     default:"5432"`
	Database string        `envconfig:"POSTGRES_DB"       required:"true"`
	User     string        `envconfig:"POSTGRES_USER"     required:"true"`
	Password string        `envconfig:"POSTGRES_PASSWORD" required:"true"`
	SSLMode  string        `envconfig:"POSTGRES_SSLMODE"  required:"true"`
	Timeout  time.Duration `envconfig:"POSTGRES_TIMEOUT"  default:"10s"`
}

func (p *Config) DSN() string {
	params := url.Values{}
	params.Set("sslmode", p.SSLMode)

	dsn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(p.User, p.Password),
		Host:     net.JoinHostPort(p.Host, p.Port),
		Path:     "/" + p.Database,
		RawPath:  "/" + url.PathEscape(p.Database),
		RawQuery: params.Encode(),
	}

	return dsn.String()
}
