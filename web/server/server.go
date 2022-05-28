package server

import (
	"net/http"

	"github.com/guibedin/poll/web/service"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	svc    service.Service
	router *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) SetService(service service.Service) {
	s.svc = service
}

func New() *Server {
	s := &Server{}
	s.routes()
	return s
}
