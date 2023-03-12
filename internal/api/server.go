package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
)

type Server struct {
	cfg            config.HTTPServer
	router         *chi.Mux
	paymentsLister internal.PaymentsLister
	paymentGetter  internal.PaymentGetter
}

func NewServer(
	cfg config.HTTPServer,
	paymentsLister internal.PaymentsLister,
	paymentGetter internal.PaymentGetter,
) *Server {
	srv := &Server{
		cfg:            cfg,
		paymentsLister: paymentsLister,
		paymentGetter:  paymentGetter,
		router:         chi.NewRouter(),
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

func (s *Server) Run(ctx context.Context) error {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Port),
		Handler:      s.router,
		IdleTimeout:  s.cfg.IdleTimeout,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(fmt.Errorf("failed to start server: %w", err))
		return err
	}

	return nil
}
