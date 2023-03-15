package asynq

import (
	"github.com/hibiken/asynq"
	"github.com/kashifsoofi/payment-gateway/internal/tasks"
	"github.com/kashifsoofi/payment-gateway/internal/tasks/asynq/handler"
	"golang.org/x/exp/slog"
)

type Server struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

func NewServer(
	cfg tasks.WorkerConfig,
	// implement asynq.Handler
	// option 1: pass all dependencies and move task handler creation here
	// option 2: pass all handlers as map[string]async.Handler
	createPaymentHandler *handler.CreatePaymentHandler,
) *Server {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.Redis.Address},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			// Queues: map[string]int{
			//     "critical": 6,
			//     "default":  3,
			//     "low":      1,
			// },
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	// mux.HandleFunc(tasks.CreatePaymentTask, tasks.HandleCreatePaymentTask)
	mux.Handle(tasks.CreatePaymentTask, createPaymentHandler)

	return &Server{
		server: server,
		mux:    mux,
	}
}

func (s *Server) Run() {
	if err := s.server.Run(s.mux); err != nil {
		slog.Error("could not run task server: %v", err)
	}
}
