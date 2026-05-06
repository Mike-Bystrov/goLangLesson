package hserver

import (
	"encoding/json"
	"net/http"
)

func (s *Server) withdrawHandler(w http.ResponseWriter, r *http.Request) {
	request := DepositRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ui, err := s.oplatiService.Withdraw(r.Context(), request.Id, request.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := NewUserResponse{
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
