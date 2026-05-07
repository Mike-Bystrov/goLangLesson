package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/thnxvlad/oplati/internal/domain"
)

func (s *Server) baseChangeBalance(w http.ResponseWriter, r *http.Request, fn func(int, int) (int, error)) {
	request := domain.ChangeBalanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ui, err := s.oplatiService.ChangeBalance(r.Context(), request.Id, request.Amount, fn)

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
