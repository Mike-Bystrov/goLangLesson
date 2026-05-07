package hserver

import (
	"net/http"

	"github.com/thnxvlad/oplati/internal/storages/inmemory"
)

func (s *Server) withdrawHandler(w http.ResponseWriter, r *http.Request) {
	s.baseChangeBalance(w, r, inmemory.MinusMoney)
}
