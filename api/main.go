package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"voting/api/poll"

	"github.com/julienschmidt/httprouter"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func logger(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		n(w, r, p)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			//r,
			time.Since(start),
		)
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
	resp := make(map[string]string)
	resp["id"] = poll.ID.Hex()
	resp["title"] = poll.Title

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func voteOnPoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Decodes body to struct
	decoder := json.NewDecoder(r.Body)
	var v poll.Vote
	err := decoder.Decode(&v)
	failOnError(err, "Failed to decode request body.")

	// Marshal vote to message
	body, err := json.Marshal(struct {
		PollId  string   `json:"pollId"`
		Options []string `json:"optionIds"`
	}{
		PollId:  ps.ByName("id"),
		Options: v.OptionIDs,
	})
	failOnError(err, "Failed to marshal vote")

	// Publish message
	err = publish(body)
	failOnError(err, "failed to publish message")

	// Write response
	resp := make(map[string]string)
	resp["message"] = "Thank you for your vote!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)
}

func getPoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pollId := ps.ByName("id")
	poll := poll.Get(pollId)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(poll)
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

	router.POST("/api/polls", logger(createPoll))
	router.GET("/api/polls/:id/", logger(getPoll))
	router.POST("/api/polls/:id/vote", logger(voteOnPoll))
	log.Fatal(http.ListenAndServe(":8080", router))
}
