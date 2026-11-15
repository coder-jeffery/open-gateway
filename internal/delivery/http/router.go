package http

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"example.com/open-gateway/internal/delivery/http/handler"
	appmw "example.com/open-gateway/internal/delivery/http/middleware"
)

// NewRouter registers routes and global middleware.
func NewRouter(logger *slog.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(chimw.RequestID)
	r.Use(chimw.Recoverer)
	r.Use(appmw.Logging(logger))

	r.Get("/healthz", handler.Health{}.ServeHTTP)

	r.Route("/v1", func(r chi.Router) {
		// LLM and proxy routes will be registered here.
	})

	return r
}