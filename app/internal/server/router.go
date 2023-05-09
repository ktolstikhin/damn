package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"

	"ktolstikhin/damn/internal/damn/vocab"
)

func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Use(httplog.RequestLogger(s.log))

	r.NotFound(s.notFound)
	r.MethodNotAllowed(s.methodNotAllowed)

	r.Get("/damn/ru", s.getDamnHandler(vocab.LanguageRU))
	r.Get("/status", s.handleGetStatus)

	return r
}
