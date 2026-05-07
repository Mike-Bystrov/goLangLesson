package hserver

import (
	"net/http"

	hmiddlewares "github.com/thnxvlad/oplati/internal/server/hmiddlewares"
	"github.com/thnxvlad/oplati/internal/service/oplati"
)

type Server struct {
	oplatiService *oplati.Service
	*http.Server
}

func NewServer(
	oplatiService *oplati.Service,
	addr string,
	mws ...func(next http.Handler) http.Handler,
) (*Server, *http.ServeMux) {
	mux := http.NewServeMux()

	httpServer := http.Server{
		Addr:    addr,
		Handler: hmiddlewares.UseMiddlewares(mux, mws),
	}

	server := &Server{
		oplatiService: oplatiService,
		Server:        &httpServer,
	}

	return server, mux
}

func NewPublicServer(
	oplatiService *oplati.Service,
	addr string,
	mws ...func(next http.Handler) http.Handler,
) *Server {
	server, mux := NewServer(oplatiService, addr, mws...)

	mux.HandleFunc("POST /transfer", server.transferHandler)
	mux.HandleFunc("POST /newUser", server.newUserHandler)
	mux.HandleFunc("POST /deposit", server.depositHandler)
	mux.HandleFunc("POST /withdraw", server.withdrawHandler)
	mux.HandleFunc("POST /getInfo/{id}", server.getInfoHandler)

	return server
}

func NewPrivateServer(
	oplatiService *oplati.Service,
	addr string,
	mws ...func(next http.Handler) http.Handler,
) *Server {
	server, mux := NewServer(oplatiService, addr, mws...)

	mux.HandleFunc("POST /getAllUsers", server.getAllUsersHandler)

	return server
}
