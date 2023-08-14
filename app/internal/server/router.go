package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"

	"ktolstikhin/damn/internal/damn"
	"ktolstikhin/damn/internal/damn/vocab"
	"ktolstikhin/damn/internal/server/handler"
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
		httprate.WithLimitHandler(handler.TooManyRequests),
	))

	r.NotFound(handler.NotFound)
	r.MethodNotAllowed(handler.MethodNotAllowed)

	ruDamner := damn.NewDamner(vocab.LanguageRU)

	r.Get("/damn/ru", handler.NewGetDamnHandler(ruDamner))
	r.Get("/status", handler.GetStatus)

	return r
}
