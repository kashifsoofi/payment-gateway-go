package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = ""

type Configuration interface{}

type Database struct {
	DatabaseURL        string `envconfig:"DATABASE_URL" required:"true"`
	LogLevel           string `envconfig:"DATABASE_LOG_LEVEL" default:"warn"`
	MaxOpenConnections int    `envconfig:"DATABASE_MAX_OPEN_CONNECTIONS" default:"10"`
}

type HTTPServer struct {
	IdleTimeout  time.Duration `envconfig:"HTTP_SERVER_IDLE_TIMEOUT" default:"60s"`
	Port         int           `envconfig:"PORT" default:"8080"`
	ReadTimeout  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"1s"`
	WriteTimeout time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"2s"`
	Store        string        `envconfig:"STORE" default:"memory"`
}

type Redis struct {
	MaxActive int    `envconfig:"REDIS_MAX_ACTIVE" default:"5"`
	MaxIdle   int    `envconfig:"REDIS_MAX_IDLE" default:"5"`
	Wait      bool   `envconfig:"REDIS_WAIT" default:"true"`
	Address   string `envconfig:"REDIS_ADDRESS" default:":6379"`
}

type TaskServer struct {
	TaskEngine string `envconfig:"DATABASE_LOG_LEVEL" default:"asynq"`
}

func Load(cfg Configuration) error {
	err := envconfig.Process(envPrefix, cfg)
	if err != nil {
		return err
	}

	return nil
}
