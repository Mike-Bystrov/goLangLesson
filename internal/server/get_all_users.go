package hserver

import (
	"encoding/json"
	"net/http"

	"github.com/thnxvlad/oplati/internal/domain"
)

type GetAllUsersResponse struct {
	Users []domain.UserInfo `json:"users"`
}

func (s *Server) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := s.oplatiService.GetAllUsers(r.Context())

	response := GetAllUsersResponse{
		Users: make([]domain.UserInfo, 0, len(users)),
	}

	for _, user := range users {
		userResponse := domain.UserInfo{
			Id:      user.Id,
			Name:    user.Name,
			Balance: user.Balance,
		}
		response.Users = append(response.Users, userResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
