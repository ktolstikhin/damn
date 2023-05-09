package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Use(s.recoverPanic)
	// TODO: use request logging middleware

	r.NotFound(s.notFound)
	r.MethodNotAllowed(s.methodNotAllowed)

	r.Get("/damn", s.handleGetDamn)

	return r
}
