package main

import (
	"context"
	"fmt"

	"github.com/kashifsoofi/payment-gateway/internal/api"
	"github.com/kashifsoofi/payment-gateway/internal/postgres"
)

func initializeServer() (*api.Server, error) {
	cfg, err := api.NewApiConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to load configuration: %w", err))
		return nil, err
	}

	store := postgres.NewPostgresStore(cfg.Database)

	server := api.NewServer(cfg.HttpServer, store, store)
	return server, nil
}

func main() {
	server, err := initializeServer()
	if err != nil {
		panic(err)
	}

	err = server.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
