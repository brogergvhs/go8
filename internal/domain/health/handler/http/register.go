package http

import (
	"github.com/go-chi/chi"

	"github.com/gmhafiz/go8/internal/domain/health"
)

func RegisterHTTPEndPoints(router *chi.Mux, uc health.UseCase) {
	h := NewHandler(uc)

	router.Route("/health", func(router chi.Router) {
		router.Get("/liveness", h.Liveness)
		router.Get("/readiness", h.Readiness)
	})
}
