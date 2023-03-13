package tasks

import (
	"context"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
)

type paymentEnqueuer struct {
	enqueuer *work.Enqueuer
}

func NewPaymentEnqueuer(
	config *config.Redis,
) *paymentEnqueuer {
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
	return &paymentEnqueuer{
		enqueuer: enqueuer,
	}
}

func (e *paymentEnqueuer) Create(ctx context.Context, payment *internal.Payment) error {
	_, err := e.enqueuer.Enqueue("create_payment", work.Q{"payment": payment})
	return err
}
