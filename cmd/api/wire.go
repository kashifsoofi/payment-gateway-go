//go:generate wire

package main

import (
	"github.com/kashifsoofi/payment-gateway/internal/api"
	"github.com/kashifsoofi/payment-gateway/internal/postgres"

	"github.com/google/wire"
)

func InitializeServer() postgres.PostgresStore {
	wire.Build(
		api.NewApiConfig,
		wire.FieldsOf(new(api.ApiConfig), "HttpServer", "Database"),
		postgres.NewPostgresStore,
		// api.NewServer,
		// wire.Bind(new(internal.PaymentsLister), new(*postgres.PostgresStore)),
		// wire.Bind(new(internal.PaymentGetter), new(*postgres.PostgresStore)),
	)
	return postgres.PostgresStore{}
}
