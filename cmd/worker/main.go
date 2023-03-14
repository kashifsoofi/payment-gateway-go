package main

import (
	"fmt"

	"github.com/kashifsoofi/payment-gateway/internal/postgres"
	"github.com/kashifsoofi/payment-gateway/internal/tasks"
	"github.com/kashifsoofi/payment-gateway/internal/tasks/handler"
)

func initializeServer() (*tasks.Server, error) {
	cfg, err := tasks.NewWorkerConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to load configuration: %w", err))
		return nil, err
	}

	store := postgres.NewPostgresStore(cfg.Database)
	createPaymentsHandler := handler.NewCreatePaymentHandler(store, store, store)
	paymentsContext := tasks.NewPaymentsContext(createPaymentsHandler)

	service := tasks.NewServer(cfg, *paymentsContext)
	return service, nil
}

func main() {
	server, err := initializeServer()
	if err != nil {
		panic(err)
	}

	server.Run()
}
