package tasks

import (
	"context"

	"github.com/gocraft/work"
	"github.com/kashifsoofi/payment-gateway/internal"
)

type CreatePaymentContext struct {
	paymentCreator internal.PaymentCreator
}

func NewCreatePaymentContext(
	paymentCreator internal.PaymentCreator,
) *CreatePaymentContext {
	return &CreatePaymentContext{
		paymentCreator: paymentCreator,
	}
}

func (c *CreatePaymentContext) CreatePaymentHandler(job *work.Job) error {
	payment := job.Args["payment"].(*internal.Payment)

	err := c.paymentCreator.Create(context.TODO(), payment)
	if err != nil {
		return err
	}

	return nil
}
