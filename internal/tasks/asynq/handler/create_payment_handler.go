package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/kashifsoofi/payment-gateway/internal"
	"golang.org/x/exp/slog"
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

func (h *CreatePaymentHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	cmd := internal.CreatePaymentCommand{}
	if err := json.Unmarshal(t.Payload(), &cmd); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	slog.Info("Create Payment Handler")

	p, err := h.paymentGetter.Get(ctx, cmd.Id)
	if err != nil {
		return fmt.Errorf("error getting payment %v", err)
	}

	if p != nil {
		return fmt.Errorf("payment with id: [%v] already exists: %w", cmd.Id, asynq.SkipRetry)
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
