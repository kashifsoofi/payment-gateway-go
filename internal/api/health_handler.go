package api

import (
	"net/http"

	"github.com/go-chi/render"
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
