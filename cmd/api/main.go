package main

import (
	"context"
	"fmt"

	"github.com/kashifsoofi/payment-gateway/internal/api"
	"github.com/kashifsoofi/payment-gateway/internal/postgres"
	"github.com/kashifsoofi/payment-gateway/internal/tasks/enqueuer"
)

func initializeServer() (*api.Server, error) {
	cfg, err := api.NewApiConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to load configuration: %w", err))
		return nil, err
	}

	store := postgres.NewPostgresStore(cfg.Database)
	enqueuer := enqueuer.NewPaymentsEnqueuer(&cfg.Redis)

	server := api.NewServer(cfg.HttpServer, store, store, enqueuer)
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
