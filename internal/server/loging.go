package hserver

import (
	"encoding/json"
	"net/http"

	auth "github.com/thnxvlad/oplati/internal/service/authorization"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type NewUserResponse struct {
	Token string `json:"token"`
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	request := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := auth.Login(request.Login, request.Password)
	if err != nil || token == "" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := NewUserResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
