package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v4emb"
)

type Server struct {
	cfg            config.HTTPServer
	handler        http.Handler
	paymentsLister internal.PaymentsLister
	paymentGetter  internal.PaymentGetter
}

func NewServer(
	cfg config.HTTPServer,
	paymentsLister internal.PaymentsLister,
	paymentGetter internal.PaymentGetter,
) *Server {
	s := web.DefaultService()
	s.OpenAPI.Info.Title = "Payment Gateway"
	s.OpenAPI.Info.WithDescription("Payment Gateway API")
	s.OpenAPI.Info.Version = "1.0.0"

	s.Get("/health", getHealthUsecase())

	s.Get("/payments/{id}", getPaymentUsecase(paymentGetter))
	// s.Route("/payments", func(r chi.Router) {
	// 	s.Get("/", getPaymentUsecase(paymentGetter))
	// 	// 	r.Post("/", s.createPayment())
	// 	// 	r.Route("/{id}", func(r chi.Router) {
	// 	// 		r.Get("/", s.getPayment())
	// 	// 	})
	// })

	// load swagger UI at the root
	s.Method(http.MethodGet, "/openapi.json", s.OpenAPICollector)
	s.Mount("/", swgui.New(s.OpenAPI.Info.Title, "/openapi.json", "/"))

	srv := &Server{
		cfg:            cfg,
		paymentsLister: paymentsLister,
		paymentGetter:  paymentGetter,
		handler:        s,
	}

	// srv.routes()

	return srv
}

// func (s *Server) routes() {
// 	s.router.Get("/health", s.getHealth())

// 	s.router.Route("/payments", func(r chi.Router) {
// 		r.Get("/", s.listPayments())
// 		r.Post("/", s.createPayment())
// 		r.Route("/{id}", func(r chi.Router) {
// 			r.Get("/", s.getPayment())
// 		})
// 	})
// }

func (s *Server) Run(ctx context.Context) error {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Port),
		Handler:      s.handler,
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
