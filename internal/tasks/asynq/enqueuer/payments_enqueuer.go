package enqueuer

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
	internal_asynq "github.com/kashifsoofi/payment-gateway/internal/tasks/asynq"
	"golang.org/x/exp/slog"
)

type paymentsEnqueuer struct {
	redisAddr string
}

func NewPaymentsEnqueuer(
	cfg config.Redis,
) *paymentsEnqueuer {
	return &paymentsEnqueuer{
		redisAddr: cfg.Address,
	}
}

func (e *paymentsEnqueuer) Enqueue(ctx context.Context, cmd *internal.CreatePaymentCommand) error {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: e.redisAddr})
	defer client.Close()

	task, err := internal_asynq.NewCreatePaymentTask(cmd)
	if err != nil {
		slog.ErrorCtx(ctx, "could not create task: %v", err)
		return err
	}

	info, err := client.Enqueue(task)
	if err != nil {
		slog.ErrorCtx(ctx, "could not enqueue task: %v", err)
		return err
	}
	slog.InfoCtx(ctx, "enqueued task: id=[%s] queue=[%s]", info.ID, info.Queue)
	return nil
}
