package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) routes() {
	s.router = httprouter.New()

	s.router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	// Polls Routes
	s.router.GET("/api/polls", s.logger(s.handlePollsGet()))
	s.router.POST("/api/polls", s.logger(s.handlePollsCreate()))
	s.router.GET("/api/polls/:id/", s.logger(s.handlePollGet()))
	s.router.POST("/api/polls/:id/vote", s.logger(s.handlePollVote()))
	s.router.POST("/api/polls/:id/votes", s.logger(s.handlePollVotes()))
}
