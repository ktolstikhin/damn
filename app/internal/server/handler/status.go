package handler

import (
	"net/http"

	"ktolstikhin/damn/internal/server/response"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	err := response.JSON(w, http.StatusOK, StatusResponse{
		Status: "ok",
	})
	if err != nil {
		serverError(w, r, err)
	}
}
