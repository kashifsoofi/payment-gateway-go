package tasks

import (
	"context"

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
	payment := job.Args["payment"].(*internal.Payment)

	err := c.createPaymentHandler.Handle(context.TODO(), payment)
	if err != nil {
		return err
	}

	return nil
}
