package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/thnxvlad/oplati/internal/domain"
)

type TransferRequest struct {
	UserIdFrom uuid.UUID `json:"id_from"`
	UserIdTo   uuid.UUID `json:"id_to"`
	Amount     int       `json:"amount"`
}

func (s *Server) Transfer(w http.ResponseWriter, r *http.Request) {
	request := TransferRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users, err := s.oplatiService.Transfer(r.Context(), request.UserIdFrom, request.UserIdTo, request.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//будет потом изменено, но пока так
	response := domain.UserInfo{
		Id:      users[1].Id,
		Name:    users[1].Name,
		Balance: users[1].Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
