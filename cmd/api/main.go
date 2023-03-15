package main

import (
	"context"
	"fmt"

	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/api"
	"github.com/kashifsoofi/payment-gateway/internal/postgres"
	asynq "github.com/kashifsoofi/payment-gateway/internal/tasks/asynq/enqueuer"
	gocraft "github.com/kashifsoofi/payment-gateway/internal/tasks/gocraft/enqueuer"
)

func initializeServer() (*api.Server, error) {
	cfg, err := api.NewApiConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to load configuration: %w", err))
		return nil, err
	}

	store := postgres.NewPostgresStore(cfg.Database)
	var createPaymentEnqueuer internal.CreatePaymentEnqueuer
	if cfg.TaskServer.TaskEngine == "asynq" {
		createPaymentEnqueuer = asynq.NewPaymentsEnqueuer(cfg.Redis)
	} else {
		createPaymentEnqueuer = gocraft.NewPaymentsEnqueuer(cfg.Redis)
	}

	server := api.NewServer(cfg.HttpServer, store, store, createPaymentEnqueuer)
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
