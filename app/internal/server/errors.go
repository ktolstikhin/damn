package server

import (
	"net/http"
	"strings"

	"ktolstikhin/damn/internal/server/response"
)

type ErrMessage struct {
	Error string `json:"error"`
}

func (s *Server) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	err := response.JSON(w, status, ErrMessage{
		Error: strings.ToLower(message),
	})
	if err != nil {
		// TODO: app.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) serverError(w http.ResponseWriter, r *http.Request, err error) {
	// TODO: app.logger.Error(err)
	s.errorMessage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.errorMessage(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func (s *Server) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	s.errorMessage(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
}

func (s *Server) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	s.errorMessage(w, r, http.StatusBadRequest, err.Error())
}
