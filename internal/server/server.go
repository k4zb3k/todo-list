package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k4zb3k/todo-list/internal/service"
)

type Server struct {
	Mux 	 *mux.Router
	Services *service.Services
}

func NewServer(mux *mux.Router, services *service.Services) *Server {
	return &Server{
		Mux: mux,
		Services: services,
	}
}

func(s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}
//=================================================================

func(s *Server) Init() {
	
	mainRout := s.Mux.PathPrefix("/api/v1").Subrouter()

	mainRout.HandleFunc("/", s.Index).Methods("POST")
}