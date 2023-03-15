package enqueuer

import (
	"context"
	"encoding/json"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
	"github.com/kashifsoofi/payment-gateway/internal/tasks"
)

type paymentsEnqueuer struct {
	enqueuer *work.Enqueuer
}

func NewPaymentsEnqueuer(
	cfg config.Redis,
) *paymentsEnqueuer {
	// Make a redis pool
	var redisPool = &redis.Pool{
		MaxActive: cfg.MaxActive,
		MaxIdle:   cfg.MaxIdle,
		Wait:      cfg.Wait,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Address)
		},
	}

	var enqueuer = work.NewEnqueuer("payment-gateway", redisPool)
	return &paymentsEnqueuer{
		enqueuer: enqueuer,
	}
}

func (e *paymentsEnqueuer) Enqueue(ctx context.Context, cmd *internal.CreatePaymentCommand) error {
	createPaymentCommandJson, _ := json.Marshal(cmd)
	_, err := e.enqueuer.EnqueueUnique(tasks.CreatePaymentTask, work.Q{"create_payment_command_json": string(createPaymentCommandJson)})
	return err
}
