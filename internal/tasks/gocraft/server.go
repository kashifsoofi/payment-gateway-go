package tasks

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/kashifsoofi/payment-gateway/internal/tasks"
)

type Server struct {
	pools []*work.WorkerPool
}

func NewServer(
	cfg tasks.WorkerConfig,
	paymentsContext PaymentsContext,
) *Server {
	// Make a redis pool
	var redisPool = &redis.Pool{
		MaxActive: cfg.Redis.MaxActive,
		MaxIdle:   cfg.Redis.MaxIdle,
		Wait:      cfg.Redis.Wait,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.Redis.Address)
		},
	}

	pools := []*work.WorkerPool{}

	paymentsPool := work.NewWorkerPool(paymentsContext, 10, "payment-gateway", redisPool)

	// Add middleware that will be executed for each job
	// pool.Middleware((*CreatePaymentContext).Log)

	// Map the name of the job to handler functions
	paymentsPool.Job(tasks.CreatePaymentTask, paymentsContext.CreatePayment)

	pools = append(pools, paymentsPool)

	return &Server{
		pools: pools,
	}
}

func (s *Server) Start() {
	for _, p := range s.pools {
		p.Start()
	}
}

func (s *Server) Stop() {
	for _, p := range s.pools {
		p.Stop()
	}
}

func (s *Server) Run() {
	s.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	s.Stop()
}
