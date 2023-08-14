package handler

import (
	"net/http"
	"strings"

	"ktolstikhin/damn/internal/server/response"
)

type ErrMessage struct {
	Error string `json:"error"`
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	errorStatus(w, r, http.StatusNotFound)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	errorStatus(w, r, http.StatusMethodNotAllowed)
}

func TooManyRequests(w http.ResponseWriter, r *http.Request) {
	errorStatus(w, r, http.StatusTooManyRequests)
}

func errorStatus(w http.ResponseWriter, r *http.Request, status int) {
	errorStatusMessage(w, r, status, strings.ToLower(http.StatusText(status)))
}

func errorStatusMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	err := response.JSON(w, status, ErrMessage{Error: message})
	if err != nil {
		log := requestLogger(r)
		log.Error().Err(err).Msg("Failed to write json response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func serverError(w http.ResponseWriter, r *http.Request, err error) {
	log := requestLogger(r)
	log.Error().Err(err).Msg("Server error")
	errorStatus(w, r, http.StatusInternalServerError)
}

func badRequest(w http.ResponseWriter, r *http.Request, err error) {
	errorStatusMessage(w, r, http.StatusBadRequest, err.Error())
}

func unprocessableEntity(w http.ResponseWriter, r *http.Request, err error) {
	errorStatusMessage(w, r, http.StatusUnprocessableEntity, err.Error())
}
