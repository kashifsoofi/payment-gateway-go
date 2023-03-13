package tasks

import (
	"os"
	"os/signal"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
)

type worker struct {
	pool *work.WorkerPool
}

func NewWorker(
	config *config.Redis,
	paymentCreator internal.PaymentCreator,
) *worker {
	// Make a redis pool
	var redisPool = &redis.Pool{
		MaxActive: config.MaxActive,
		MaxIdle:   config.MaxIdle,
		Wait:      config.Wait,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Address)
		},
	}

	context := NewCreatePaymentContext(paymentCreator)
	pool := work.NewWorkerPool(context, 10, "payment-gateway", redisPool)

	// Add middleware that will be executed for each job
	// pool.Middleware((*CreatePaymentContext).Log)

	// Map the name of the job to handler functions
	pool.Job("create_payment", context.CreatePaymentHandler)

	return &worker{
		pool: pool,
	}
}

func (w *worker) Start() {
	// Start processing jobs
	w.pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	w.pool.Stop()
}
