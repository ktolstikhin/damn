package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"

	"ktolstikhin/damn/internal/damn/vocab"
)

func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Use(httplog.RequestLogger(s.log))
	r.Use(httprate.Limit(
		2,
		time.Second,
		httprate.WithKeyFuncs(
			httprate.KeyByIP,
			httprate.KeyByEndpoint,
		),
		httprate.WithLimitHandler(s.tooManyRequests),
	))

	r.NotFound(s.notFound)
	r.MethodNotAllowed(s.methodNotAllowed)

	r.Get("/damn/ru", s.getDamnHandler(vocab.LanguageRU))
	r.Get("/status", s.handleGetStatus)

	return r
}
