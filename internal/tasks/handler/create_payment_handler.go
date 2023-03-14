package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/kashifsoofi/payment-gateway/internal"
)

type CreatePaymentHandler struct {
	paymentGetter        internal.PaymentGetter
	paymentCreator       internal.PaymentCreator
	paymentStatusUpdater internal.PaymentStatusUpdater
}

func NewCreatePaymentHandler(
	paymentGetter internal.PaymentGetter,
	paymentCreator internal.PaymentCreator,
	paymentStatusUpdater internal.PaymentStatusUpdater,
) *CreatePaymentHandler {
	return &CreatePaymentHandler{
		paymentGetter:        paymentGetter,
		paymentCreator:       paymentCreator,
		paymentStatusUpdater: paymentStatusUpdater,
	}
}

func (h *CreatePaymentHandler) Handle(ctx context.Context, cmd *internal.CreatePaymentCommand) error {
	p, err := h.paymentGetter.Get(ctx, cmd.Id)
	if err != nil {
		return err
	}

	if p != nil {
		return fmt.Errorf("payment with id: %v already exists", cmd.Id)
	}

	status := internal.PaymentStatusApproved // TODO get this by calling payments system
	payment := internal.NewPayment(
		cmd.Id, cmd.CardHolderName, cmd.CardNumber, cmd.ExpiryMonth,
		cmd.ExpiryYear, cmd.Amount, cmd.CurrencyCode, cmd.Reference,
		status, time.Now().UTC(), time.Now().UTC(),
	)

	err = h.paymentCreator.Create(ctx, payment)
	if err != nil {
		return err
	}

	// TODO publish event

	return nil
}
