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
	Words []string `json:"words"`
}

func (s *Server) getDamnHandler(lang vocab.Language) http.HandlerFunc {
	damner := damn.NewDamner(lang)

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			level int
			err   error
			opts  []vocab.Option
		)

		if r.URL.Query().Has("level") {
			levelStr := r.URL.Query().Get("level")
			level, err = strconv.Atoi(levelStr)
			if err != nil {
				s.unprocessableEntity(w, r, fmt.Errorf("invalid query level: %s", levelStr))

				return
			}
		}

		if r.URL.Query().Has("gender") {
			genderStr := r.URL.Query().Get("gender")
			gender, ok := vocab.StrToGenderMap[genderStr]
			if !ok {
				s.unprocessableEntity(w, r, fmt.Errorf("invalid query gender: %s", genderStr))

				return
			}
			opts = append(opts, vocab.WithGender(gender))
		}

		if r.URL.Query().Has("obscene") {
			obsceneStr := r.URL.Query().Get("obscene")
			obscene, err := strconv.ParseBool(obsceneStr)
			if err != nil {
				s.unprocessableEntity(w, r, fmt.Errorf("invalid query obscene: %s", obsceneStr))

				return
			}
			opts = append(opts, vocab.WithObscene(obscene))
		}

		words := damner.DamnYou(level, opts...)

		err = response.JSON(w, http.StatusOK, DamnResponse{
			Words: words,
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
