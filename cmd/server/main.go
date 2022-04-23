package main

import (
	"log"
	"net/http"

	"github.com/guibedin/poll"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	router.POST("/api/polls", poll.CreatePoll)
	router.GET("/api/polls", poll.GetPolls)
	router.GET("/api/polls/:id/", poll.GetPoll)
	router.POST("/api/polls/:id/vote", poll.VoteOnPoll)

	log.Fatal(http.ListenAndServe(":8080", router))
}
