package service

import (
	"fmt"

	"github.com/guibedin/poll/web/domain"
)

type voteServiceError struct {
	message string
}

func (e voteServiceError) Error() string {
	return fmt.Sprintf("VoteServicceError: %s", e.message)
}

// GetVoteCountByOptionID returns the number of votes of a given option
func (s *Service) GetVoteCountByOptionID(id int) (int, error) {
	votes, err := s.repo.GetVoteCountByOptionID(id)
	if err != nil {
		panic(err)
	}
	return votes, nil
}

func (s *Service) GetVotesByOptionID(id int) ([]domain.Vote, error) {
	votes, err := s.repo.GetVotesByOptionID(id)
	if err != nil {
		panic(err)
	}
	return votes, nil
}

// AddVote publishes the vote to the queue
func (s *Service) AddVote(v domain.Vote) error {
	poll, err := s.repo.GetPoll(v.PollId)
	if err != nil {
		panic(err)
	}

	for _, option := range poll.Options {
		if v.OptionId == option.ID {
			return s.repo.AddVote(v)
		}
	}

	return voteServiceError{"Option must belong to selected Poll"}
}

// AddVotes publishes the votes to the queue
func (s *Service) AddVotes(v []domain.Vote) error {

	// check if all pollIds are the same
	var pollId int
	if len(v) > 0 {
		pollId = v[0].PollId
	} else {
		return voteServiceError{"Empty list of votes"}
	}
	for _, vote := range v {
		if vote.PollId != pollId {
			return voteServiceError{"Votes must belong to the same Poll"}
		}
	}

	// check if all voted options belong to poll
	poll, err := s.repo.GetPoll(pollId)
	if err != nil {
		panic(err)
	}
	for _, vote := range v {
		if !optionsContains(poll.Options, vote.OptionId) {
			return voteServiceError{"All Options must belong to selected Poll"}
		}
	}

	// adds all votes to repository
	err = s.repo.AddVotes(v)
	if err != nil {
		return err
	}

	return nil
}

func optionsContains(options []domain.Option, id int) bool {
	for _, o := range options {
		if o.ID == id {
			return true
		}
	}
	return false
}
