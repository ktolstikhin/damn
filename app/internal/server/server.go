package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

const (
	idleTimeout  = 2 * time.Minute
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

type Server struct {
	addr string
	log  zerolog.Logger
}

func New(addr string, log zerolog.Logger) *Server {
	return &Server{
		addr: addr,
		log:  log,
	}
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.addr,
		Handler:      s.router(),
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	shutdownError := make(chan error)

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return <-shutdownError
}

func (s *Server) requestLogger(r *http.Request) zerolog.Logger {
	return httplog.LogEntry(r.Context())
}
