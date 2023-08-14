package handler

import (
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

func requestLogger(r *http.Request) zerolog.Logger {
	return httplog.LogEntry(r.Context())
}
