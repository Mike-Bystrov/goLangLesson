package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/thnxvlad/oplati/internal/domain"
)

func (s *Server) depositHandler(w http.ResponseWriter, r *http.Request) {
	request := domain.ChangeBalanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ui, err := s.oplatiService.Deposit(r.Context(), request.Id, request.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := domain.UserInfo{
		Id:      ui.Id,
		Name:    ui.Name,
		Balance: ui.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}