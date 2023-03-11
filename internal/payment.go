package internal

import (
	"github.com/google/uuid"
)

type Payment struct {
	Id             uuid.UUID     `json:"id"`
	MerchantId     uuid.UUID     `json:"merchant_id"`
	CardHolderName string        `json:"card_holder_name"`
	CardNumber     string        `json:"card_number"`
	ExpiryMonth    int           `json:"expiry_month"`
	ExpiryYear     int           `json:"expiry_year"`
	Amount         float64       `json:"amount"`
	CurrencyCode   string        `json:"currency_code"`
	Reference      string        `json:"reference"`
	Status         PaymentStatus `json:"status"`
	CreatedAt      string        `json:"created_at"`
	UpdatedAt      string        `json:"updated_at"`
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
