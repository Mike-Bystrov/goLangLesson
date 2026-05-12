package hserver

import (
	"encoding/json"
	"net/http"

	auth "github.com/thnxvlad/oplati/internal/service/authorization"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	req := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err1 := s.oplatiService.CreateUser(r.Context(), req.Name)

	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	err = auth.SignUp(req.Login, req.Password, user.Id.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := auth.Login(req.Login, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := RegisterResponse{
		Token: token,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
