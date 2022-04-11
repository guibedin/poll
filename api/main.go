package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"voting/api/poll"

	"github.com/julienschmidt/httprouter"
)

type vote struct {
	Poll   string `json:"poll"`
	Option string `json:"option"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createPoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decode poll
	decoder := json.NewDecoder(r.Body)
	var p poll.Poll
	err := decoder.Decode(&p)
	failOnError(err, "Failed to decode poll")

	// Create poll and save to DB
	poll, err := poll.New(p.Title, p.Options)
	failOnError(err, "Failed to create poll")

	// Return poll ID
	fmt.Println(poll)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["title"] = poll.Title
	json.NewEncoder(w).Encode(resp)
}

func votes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Decodes body to struct
	decoder := json.NewDecoder(r.Body)
	var v vote
	err := decoder.Decode(&v)
	failOnError(err, "Failed to decode request body.")

	// Publish message
	//body := v.Option

	// Write response
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Thank you for your vote!"
	json.NewEncoder(w).Encode(resp)
}

func results(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Write response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["results"] = "Option 1: 1; Option 2: 2"
	json.NewEncoder(w).Encode(resp)
}

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

	router.POST("/api/polls", createPoll)
	router.POST("/api/polls/:id/vote", votes)
	router.GET("/api/polls/:id/results", results)
	log.Fatal(http.ListenAndServe(":8080", router))
}
