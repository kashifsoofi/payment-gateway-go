package main

import (
	"fmt"

	"github.com/kashifsoofi/payment-gateway/internal/postgres"
	"github.com/kashifsoofi/payment-gateway/internal/tasks"
	asynq "github.com/kashifsoofi/payment-gateway/internal/tasks/asynq"
	asynq_handler "github.com/kashifsoofi/payment-gateway/internal/tasks/asynq/handler"
	gocraft "github.com/kashifsoofi/payment-gateway/internal/tasks/gocraft"
	gocraft_handler "github.com/kashifsoofi/payment-gateway/internal/tasks/gocraft/handler"
)

func initializeServer() (tasks.Server, error) {
	cfg, err := tasks.NewWorkerConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to load configuration: %w", err))
		return nil, err
	}

	store := postgres.NewPostgresStore(cfg.Database)

	var server tasks.Server
	if cfg.TaskServer.TaskEngine == "asynq" {
		createPaymentsHandler := asynq_handler.NewCreatePaymentHandler(store, store, store)

		server = asynq.NewServer(cfg, createPaymentsHandler)
	} else {
		createPaymentsHandler := gocraft_handler.NewCreatePaymentHandler(store, store, store)
		paymentsContext := gocraft.NewPaymentsContext(createPaymentsHandler)

		server = gocraft.NewServer(cfg, *paymentsContext)
	}

	return server, nil
}

func main() {
	server, err := initializeServer()
	if err != nil {
		panic(err)
	}

	server.Run()
}
