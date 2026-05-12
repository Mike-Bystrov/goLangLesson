package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/thnxvlad/oplati/internal/domain"
	"github.com/thnxvlad/oplati/internal/server/hmiddlewares"
)

func (s *Server) depositHandler(w http.ResponseWriter, r *http.Request) {
	request := domain.ChangeBalanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
	}
	
	accountId := r.Context().Value(hmiddlewares.AccountIdContextKey{}).(string)
	
	if accountId == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ui, err := s.oplatiService.Deposit(r.Context(), uuid.MustParse(accountId), request.Amount)

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