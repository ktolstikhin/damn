package server

import (
	"errors"
	"net/http"
	"strings"

	"ktolstikhin/damn/internal/server/response"
)

var ErrTooManyRequests = errors.New("too many requests")

type ErrMessage struct {
	Error string `json:"error"`
}

func (s *Server) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	err := response.JSON(w, status, ErrMessage{Error: message})
	if err != nil {
		log := s.RequestLogger(r)
		log.Error().Err(err).Msg("Failed to write json response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) serverError(w http.ResponseWriter, r *http.Request, err error) {
	log := s.RequestLogger(r)
	log.Error().Err(err).Msg("Server error")
	status := http.StatusInternalServerError
	message := strings.ToLower(http.StatusText(status))
	s.errorMessage(w, r, status, message)
}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	status := http.StatusNotFound
	message := strings.ToLower(http.StatusText(status))
	s.errorMessage(w, r, status, message)
}

func (s *Server) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	status := http.StatusMethodNotAllowed
	message := strings.ToLower(http.StatusText(status))
	s.errorMessage(w, r, status, message)
}

func (s *Server) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	s.errorMessage(w, r, http.StatusBadRequest, err.Error())
}

func (s *Server) unprocessableEntity(w http.ResponseWriter, r *http.Request, err error) {
	s.errorMessage(w, r, http.StatusUnprocessableEntity, err.Error())
}

func (s *Server) tooManyRequests(w http.ResponseWriter, r *http.Request, err error) {
	s.errorMessage(w, r, http.StatusTooManyRequests, err.Error())
}
