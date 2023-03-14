package tasks

import (
	"github.com/kashifsoofi/payment-gateway/internal/config"
)

type WorkerConfig struct {
	Database config.Database
	Redis    config.Redis
}

func NewWorkerConfig() (WorkerConfig, error) {
	var cfg WorkerConfig
	if err := config.Load(&cfg); err != nil {
		return WorkerConfig{}, err
	}

	return cfg, nil
}
