package asynq

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/tasks"
)

func NewCreatePaymentTask(cmd *internal.CreatePaymentCommand) (*asynq.Task, error) {
	payload, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(tasks.CreatePaymentTask, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}
