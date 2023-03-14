package enqueuer

import (
	"context"

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
	config *config.Redis,
) *paymentsEnqueuer {
	// Make a redis pool
	var redisPool = &redis.Pool{
		MaxActive: config.MaxActive,
		MaxIdle:   config.MaxIdle,
		Wait:      config.Wait,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Address)
		},
	}

	var enqueuer = work.NewEnqueuer("payment-gateway", redisPool)
	return &paymentsEnqueuer{
		enqueuer: enqueuer,
	}
}

func (e *paymentsEnqueuer) Create(ctx context.Context, payment *internal.Payment) error {
	_, err := e.enqueuer.EnqueueUnique(tasks.CreatePaymentTask, work.Q{"payment": payment})
	return err
}
