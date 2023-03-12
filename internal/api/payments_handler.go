package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/kashifsoofi/payment-gateway/internal"
)

type paymentResponse struct {
	Id             uuid.UUID              `json:"id"`
	MerchantId     uuid.UUID              `json:"merchant_id"`
	CardHolderName string                 `json:"card_holder_name"`
	CardNumber     string                 `json:"card_number"`
	ExpiryMonth    int                    `json:"expiry_month"`
	ExpiryYear     int                    `json:"expiry_year"`
	Amount         float64                `json:"amount"`
	CurrencyCode   string                 `json:"currency_code"`
	Reference      string                 `json:"reference"`
	Status         internal.PaymentStatus `json:"status"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

func NewPaymentResponse(p *internal.Payment) *paymentResponse {
	return &paymentResponse{
		Id:             p.Id,
		MerchantId:     p.MerchantId,
		CardHolderName: p.CardHolderName,
		CardNumber:     p.CardNumber,
		ExpiryMonth:    p.ExpiryMonth,
		ExpiryYear:     p.ExpiryYear,
		Amount:         p.Amount,
		CurrencyCode:   p.CurrencyCode,
		Reference:      p.Reference,
		Status:         p.Status,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func (pr *paymentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPaymentsListResponse(payments []*internal.Payment) []render.Renderer {
	list := []render.Renderer{}
	for _, payment := range payments {
		mr := NewPaymentResponse(payment)
		list = append(list, mr)
	}
	return list
}

func (s *Server) listPayments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		merchantIdHeader, ok := r.Header["Merchant-Id"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}
		merchantId, err := uuid.Parse(merchantIdHeader[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		payments, err := s.merchantPaymentsGetter.Get(merchantId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		render.RenderList(w, r, NewPaymentsListResponse(payments))
	}
}

func (s *Server) createPayment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *Server) getPayment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		payment, err := s.paymentGetter.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		pr := NewPaymentResponse(payment)
		render.Render(w, r, pr)
	}
}
