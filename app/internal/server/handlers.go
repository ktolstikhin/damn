package server

import (
	"fmt"
	"net/http"
	"strconv"

	"ktolstikhin/damn/internal/damn"
	"ktolstikhin/damn/internal/damn/vocab"
	"ktolstikhin/damn/internal/server/response"
)

type DamnResponse struct {
	Tokens  []string     `json:"tokens"`
	Gender  vocab.Gender `json:"gender"`
	Obscene bool         `json:"obscene"`
}

func (s *Server) getDamnHandler(lang vocab.Language) http.HandlerFunc {
	damner := damn.NewDamner(lang)

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			level   = 1
			gender  = vocab.GenderMasculine
			obscene = false
			err     error
		)

		if r.URL.Query().Has("level") {
			levelStr := r.URL.Query().Get("level")
			level, err = strconv.Atoi(levelStr)
			if err != nil {
				s.unprocessableEntity(w, r, fmt.Errorf("invalid level: %s", levelStr))

				return
			}
		}

		if r.URL.Query().Has("gender") {
			genderStr := r.URL.Query().Get("gender")
			gender, err = vocab.ParseGender(genderStr)
			if err != nil {
				s.unprocessableEntity(w, r, fmt.Errorf("invalid gender: %s", genderStr))

				return
			}
		}

		if r.URL.Query().Has("obscene") {
			obsceneStr := r.URL.Query().Get("obscene")
			obscene, err = strconv.ParseBool(obsceneStr)
			if err != nil {
				s.unprocessableEntity(w, r, fmt.Errorf("invalid obscene: %s", obsceneStr))

				return
			}
		}

		tokens := damner.DamnYou(
			level,
			vocab.WithGender(gender),
			vocab.WithObscene(obscene),
		)

		err = response.JSON(w, http.StatusOK, DamnResponse{
			Tokens:  tokens,
			Gender:  gender,
			Obscene: obscene,
		})
		if err != nil {
			s.serverError(w, r, err)
		}
	}
}

type StatusResponse struct {
	Status string `json:"status"`
}

func (s *Server) handleGetStatus(w http.ResponseWriter, r *http.Request) {
	err := response.JSON(w, http.StatusOK, StatusResponse{
		Status: "ok",
	})
	if err != nil {
		s.serverError(w, r, err)
	}
}
