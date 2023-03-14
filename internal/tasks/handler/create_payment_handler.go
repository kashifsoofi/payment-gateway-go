package handler

import (
	"context"
	"fmt"

	"github.com/kashifsoofi/payment-gateway/internal"
)

type CreatePaymentHandler struct {
	paymentGetter  internal.PaymentGetter
	paymentCreator internal.PaymentCreator
}

func NewCreatePaymentHandler(
	paymentGetter internal.PaymentGetter,
	paymentCreator internal.PaymentCreator,
) *CreatePaymentHandler {
	return &CreatePaymentHandler{
		paymentGetter:  paymentGetter,
		paymentCreator: paymentCreator,
	}
}

func (h *CreatePaymentHandler) Handle(ctx context.Context, payment *internal.Payment) error {
	p, err := h.paymentGetter.Get(ctx, payment.Id)
	if err != nil {
		return nil
	}

	if p != nil {
		return fmt.Errorf("payment with id: %v already exists", payment.Id)
	}

	err = h.paymentCreator.Create(ctx, payment)
	if err != nil {
		return err
	}

	// TODO publish event

	return nil
}
