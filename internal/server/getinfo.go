package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type GetInfoRequest struct {
	Id uuid.UUID `json:"id"`
}

func (s *Server) getInfoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ui, err := s.oplatiService.GetUser(r.Context(), id)

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
