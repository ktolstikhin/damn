package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Use(s.recoverPanic)
	r.Use(httplog.RequestLogger(s.log))

	r.NotFound(s.notFound)
	r.MethodNotAllowed(s.methodNotAllowed)

	r.Get("/damn", s.handleGetDamn)

	return r
}
