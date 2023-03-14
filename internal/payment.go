package internal

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	Id             uuid.UUID
	MerchantId     uuid.UUID
	CardHolderName string
	CardNumber     string
	ExpiryMonth    int
	ExpiryYear     int
	Amount         float64
	CurrencyCode   string
	Reference      string
	Status         PaymentStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewPayment(
	id uuid.UUID,
	cardHolderName string,
	cardNumber string,
	expiryMonth int,
	expiryYear int,
	amount float64,
	currencyCode string,
	reference string,
	status PaymentStatus,
	createdAt time.Time,
	updatedAt time.Time,
) *Payment {
	return &Payment{
		Id:             id,
		CardHolderName: cardHolderName,
		CardNumber:     cardNumber,
		ExpiryMonth:    expiryMonth,
		ExpiryYear:     expiryYear,
		Amount:         amount,
		CurrencyCode:   currencyCode,
		Reference:      reference,
		Status:         status,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

type PaymentGetter interface {
	Get(ctx context.Context, id uuid.UUID) (*Payment, error)
}

type PaymentsLister interface {
	List(ctx context.Context, merchantId uuid.UUID) ([]*Payment, error)
}

type PaymentCreator interface {
	Create(ctx context.Context, payment *Payment) error
}

type CreatePaymentCommand struct {
	Id             uuid.UUID `json:"id"`
	CardHolderName string    `json:"card_holder_name"`
	CardNumber     string    `json:"card_number"`
	ExpiryMonth    int       `json:"expiry_month"`
	ExpiryYear     int       `json:"expiry_year"`
	Amount         float64   `json:"amount"`
	CurrencyCode   string    `json:"currency_code"`
	Reference      string    `json:"reference"`
}

func NewCreatePaymentCommand(
	id uuid.UUID,
	cardHolderName string,
	cardNumber string,
	expiryMonth int,
	expiryYear int,
	amount float64,
	currencyCode string,
	reference string,
) *CreatePaymentCommand {
	return &CreatePaymentCommand{
		Id:             id,
		CardHolderName: cardHolderName,
		CardNumber:     cardNumber,
		ExpiryMonth:    expiryMonth,
		ExpiryYear:     expiryYear,
		Amount:         amount,
		CurrencyCode:   currencyCode,
		Reference:      reference,
	}
}

type CreatePaymentEnqueuer interface {
	Enqueue(ctx context.Context, cmd *CreatePaymentCommand) error
}

type PaymentStatusUpdater interface {
	UpdatePaymentStatus(ctx context.Context, id uuid.UUID, status PaymentStatus) error
}
