package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/thnxvlad/oplati/internal/server/hmiddlewares"
)

type TransferRequest struct {
	UserIdTo   uuid.UUID `json:"id_to"`
	Amount     int       `json:"amount"`
}

func (s *Server) transferHandler(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value(hmiddlewares.AccountIdContextKey{}).(string)
	if accountId == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	request := TransferRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.oplatiService.Transfer(r.Context(), uuid.MustParse(accountId), request.UserIdTo, request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
