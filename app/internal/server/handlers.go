package server

import (
	"net/http"

	"ktolstikhin/damn/internal/server/response"
)

type DamnMessage struct {
	Damn []string `json:"damn"`
}

func (s *Server) handleGetDamn(w http.ResponseWriter, r *http.Request) {
	// TODO: call damner.DamnYou
	err := response.JSON(w, http.StatusOK, DamnMessage{})
	if err != nil {
		s.serverError(w, r, err)
	}
}
