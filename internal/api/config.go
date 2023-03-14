package api

import (
	"github.com/kashifsoofi/payment-gateway/internal/config"
)

type ApiConfig struct {
	HttpServer config.HTTPServer
	Database   config.Database
	Redis      config.Redis
}

func NewApiConfig() (ApiConfig, error) {
	var cfg ApiConfig
	if err := config.Load(&cfg); err != nil {
		return ApiConfig{}, err
	}

	return cfg, nil
}
