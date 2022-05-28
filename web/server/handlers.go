package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guibedin/poll/web/domain"
	"github.com/julienschmidt/httprouter"
)

func (s *Server) handlePollsGet() func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Get all polls from DB
		polls, err := s.svc.GetPolls()
		if err != nil {
			panic(err)
		}

		// Return list of polls
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(polls)
	}
}

func (s *Server) handlePollsCreate() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Create new poll from r.Body
		var p domain.Poll
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			panic(err)
		}

		// Save poll to DB
		id := s.svc.AddPoll(p)

		// Return pollId
		response := make(map[string]int)
		response["poll_id"] = id
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) handlePollGet() func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get poll from DB based on ID from request params
		pollId, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			panic(err)
		}

		// Return poll
		poll, err := s.svc.GetPoll(pollId)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(poll)
	}
}

func (s *Server) handlePollVote() func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pollId, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			panic(err)
		}

		var vote domain.Vote
		err = json.NewDecoder(r.Body).Decode(&vote)
		if err != nil {
			panic(err)
		}
		vote.PollId = pollId

		err = s.svc.AddVote(vote)
		if err != nil {
			panic(err)
		}

		// Return 202 accepted
		w.WriteHeader(http.StatusAccepted)
	}
}

// handlePollVotes handles Votes for multiple_choice Polls
func (s *Server) handlePollVotes() func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Receives poll_id from path var and optionIds from r.Body
		pollId, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			panic(err)
		}

		type newVote struct {
			Voter     string `json:"voter"`
			OptionIds []int  `json:"option_ids"`
		}
		var nv newVote

		err = json.NewDecoder(r.Body).Decode(&nv)
		if err != nil {
			panic(err)
		}

		// Build list of votes based on all options received
		var votes []domain.Vote
		for _, option := range nv.OptionIds {
			votes = append(votes,
				domain.Vote{
					OptionId: option,
					PollId:   pollId,
					Voter:    nv.Voter})
		}
		err = s.svc.AddVotes(votes)
		if err != nil {
			panic(err)
		}

		// Return 202 accepted
		w.WriteHeader(http.StatusAccepted)
	}
}
