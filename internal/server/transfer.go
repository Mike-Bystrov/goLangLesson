package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type TransferRequest struct {
	UserIdFrom uuid.UUID `json:"id_from"`
	UserIdTo   uuid.UUID `json:"id_to"`
	Amount     int       `json:"amount"`
}

func (s *Server) transferHandler(w http.ResponseWriter, r *http.Request) {
	request := TransferRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.oplatiService.Transfer(r.Context(), request.UserIdFrom, request.UserIdTo, request.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
