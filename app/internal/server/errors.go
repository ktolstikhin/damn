package server

import (
	"net/http"
	"strings"

	"ktolstikhin/damn/internal/server/response"
)

type ErrMessage struct {
	Error string `json:"error"`
}

func (s *Server) errorStatusMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	err := response.JSON(w, status, ErrMessage{Error: message})
	if err != nil {
		log := s.RequestLogger(r)
		log.Error().Err(err).Msg("Failed to write json response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) errorStatus(w http.ResponseWriter, r *http.Request, status int) {
	s.errorStatusMessage(w, r, status, strings.ToLower(http.StatusText(status)))
}

func (s *Server) serverError(w http.ResponseWriter, r *http.Request, err error) {
	log := s.RequestLogger(r)
	log.Error().Err(err).Msg("Server error")
	s.errorStatus(w, r, http.StatusInternalServerError)
}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.errorStatus(w, r, http.StatusNotFound)
}

func (s *Server) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	s.errorStatus(w, r, http.StatusMethodNotAllowed)
}

func (s *Server) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	s.errorStatusMessage(w, r, http.StatusBadRequest, err.Error())
}

func (s *Server) unprocessableEntity(w http.ResponseWriter, r *http.Request, err error) {
	s.errorStatusMessage(w, r, http.StatusUnprocessableEntity, err.Error())
}

func (s *Server) tooManyRequests(w http.ResponseWriter, r *http.Request) {
	s.errorStatus(w, r, http.StatusTooManyRequests)
}
