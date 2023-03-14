package tasks

import (
	"context"
	"encoding/json"

	"github.com/gocraft/work"
	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/tasks/handler"
)

type PaymentsContext struct {
	createPaymentHandler *handler.CreatePaymentHandler
}

func NewPaymentsContext(
	createPaymentHandler *handler.CreatePaymentHandler,
) *PaymentsContext {
	return &PaymentsContext{
		createPaymentHandler: createPaymentHandler,
	}
}

func (c *PaymentsContext) CreatePayment(job *work.Job) error {
	createPaymentCommandJson := job.ArgString("create_payment_command_json")
	cmd := internal.CreatePaymentCommand{}
	err := json.Unmarshal([]byte(createPaymentCommandJson), &cmd)
	if err != nil {
		return err
	}

	err = c.createPaymentHandler.Handle(context.TODO(), &cmd)
	if err != nil {
		return err
	}

	return nil
}
