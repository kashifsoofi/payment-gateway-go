package internal

import (
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

type PaymentGetter interface {
	Get(id uuid.UUID) (*Payment, error)
}

type MerchantPaymentsGetter interface {
	Get(merchantId uuid.UUID) ([]*Payment, error)
}

type PaymentCreator interface {
	Create(payment *Payment) (*Payment, error)
}
