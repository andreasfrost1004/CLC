package http

import (
	"net/http"

	"CLC/internal/http/handlers"

	"github.com/go-chi/chi/v5"
)

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/hello", handlers.Hello)
	r.Get("/item", handlers.GetItemWowhead())
	return r
}
