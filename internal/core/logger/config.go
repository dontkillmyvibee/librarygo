package logger

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}
