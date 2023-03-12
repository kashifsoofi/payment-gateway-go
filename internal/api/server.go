package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
)

type Server struct {
	cfg                    config.HTTPServer
	router                 *chi.Mux
	merchantPaymentsGetter internal.MerchantPaymentsGetter
	paymentGetter          internal.PaymentGetter
}

func NewServer(
	cfg config.HTTPServer,
	merchantPaymentsGetter internal.MerchantPaymentsGetter,
	paymentGetter internal.PaymentGetter,
) *Server {
	srv := &Server{
		cfg:                    cfg,
		merchantPaymentsGetter: merchantPaymentsGetter,
		paymentGetter:          paymentGetter,
		router:                 chi.NewRouter(),
	}

	srv.routes()

	return srv
}

func (s *Server) routes() {
	s.router.Get("/health", s.getHealth())

	s.router.Route("/payments", func(r chi.Router) {
		r.Get("/", s.listPayments())
		r.Post("/", s.createPayment())
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.getPayment())
		})
	})
}
