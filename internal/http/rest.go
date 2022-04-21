package http

import (
	"encoding/json"
	"net/http"

	"github.com/guibedin/voting/internal/create"
	"github.com/julienschmidt/httprouter"
)

func Handler(c create.Service, r read.Service, u update.Service) http.Handler {
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

	// router.POST("/api/polls", logger(createPoll))
	// router.GET("/api/polls/:id/", logger(getPoll))
	// router.POST("/api/polls/:id/vote", logger(voteOnPoll))

	return router
}

// addPoll returns a handler for POST /api/polls
func createPoll(c create.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newPoll := create.Poll
		err := decoder.Decode(&newPoll)
		if err != nil {
			log.Fatalf("%s: %s", "Failed to decode body", err)
		}

		c.CreatePoll(newPoll)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New poll created.")
	}
	
	// // Decode poll
	// decoder := json.NewDecoder(r.Body)
	// var p poll.Poll
	// err := decoder.Decode(&p)
	// failOnError(err, "Failed to decode poll")

	// // Create poll and save to DB
	// poll := poll.New(p.Title, p.Options)
	// err = poll.Save()
	// failOnError(err, "Failed to save poll")

	// // Return poll ID
	// resp := make(map[string]string)
	// resp["id"] = poll.ID.Hex()
	// resp["title"] = poll.Title

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(resp)
}

// addPoll returns a handler for POST /api/polls
func getPoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pollId := ps.ByName("id")
	poll := poll.Get(pollId)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(poll)
}

// // addPoll returns a handler for POST /api/polls
// func voteOnPoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	// Decodes body to struct
// 	decoder := json.NewDecoder(r.Body)
// 	var v poll.Vote
// 	err := decoder.Decode(&v)
// 	failOnError(err, "Failed to decode request body.")

// 	// Marshal vote to message
// 	body, err := json.Marshal(struct {
// 		PollId  string   `json:"pollId"`
// 		Options []string `json:"optionIds"`
// 	}{
// 		PollId:  ps.ByName("id"),
// 		Options: v.OptionIDs,
// 	})
// 	failOnError(err, "Failed to marshal vote")

// 	// Publish message
// 	err = publish(body)
// 	failOnError(err, "failed to publish message")

// 	// Write response
// 	resp := make(map[string]string)
// 	resp["message"] = "Thank you for your vote!"
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)
// 	json.NewEncoder(w).Encode(resp)
// }
