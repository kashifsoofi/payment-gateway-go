package api

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/swaggest/usecase"
)

type healthResponse struct {
	Ok bool `json:"ok"`
}

func (hr healthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) getHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := healthResponse{Ok: true}

		render.Render(w, r, health)
	}
}

func getHealthUsecase() usecase.IOInteractor {
	u := usecase.NewIOI(nil, new(healthResponse), func(ctx context.Context, input, output interface{}) error {
		var (
			out = output.(*healthResponse)
		)

		out.Ok = true

		return nil
	})

	u.SetDescription("Get health status")

	return u
}
