package main

import (
	"log"
	"net/http"

	"github.com/guibedin/voting"
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

	router.POST("/api/polls", voting.CreatePoll)
	router.GET("/api/polls", voting.GetPolls)
	router.GET("/api/polls/:id/", voting.GetPoll)
	router.POST("/api/polls/:id/vote", voting.VoteOnPoll)

	log.Fatal(http.ListenAndServe(":8080", router))
}
