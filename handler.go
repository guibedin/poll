package voting

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func CreatePoll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Create new poll from r.Body
	var jPoll JsonPoll
	err := json.NewDecoder(r.Body).Decode(&jPoll)
	if err != nil {
		panic(err)
	}

	// Save poll to DB
	id := jPoll.Save()

	// Return pollId
	response := make(map[string]int)
	response["poll_id"] = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func VoteOnPoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Receives poll_id from path var and optionIds from r.Body
	pollId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		panic(err)
	}
	var jOptionIds JsonOptionIds
	err = json.NewDecoder(r.Body).Decode(&jOptionIds)
	if err != nil {
		panic(err)
	}

	// Build message body and publish to 'votes' queue
	mqVote := MqVote{
		pollId,
		jOptionIds.OptionIds,
	}
	body, err := json.Marshal(mqVote)
	if err != nil {
		panic(err)
	}
	send(body)

	// Return 202 accepted
	w.WriteHeader(http.StatusAccepted)
}

func GetPoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get poll from DB based on ID from request params
	pollId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		panic(err)
	}

	// Return poll
	poll, err := Get(pollId)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(poll)
}

func getPolls(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all polls from DB

		// Return list of polls
	}
}
