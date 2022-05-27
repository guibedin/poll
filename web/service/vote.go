package service

import (
	"github.com/guibedin/poll/web/domain"
)

// GetVoteCountByOptionID returns the number of votes of a given option
func (s *Service) GetVoteCountByOptionID(id int) (int, error) {
	votes, err := s.repo.GetVoteCountByOptionID(id)
	if err != nil {
		panic(err)
	}
	return votes, nil
}

func (s *Service) GetVote() (domain.Vote, error) {
	vote, err := s.repo.GetVote()
	if err != nil {
		panic(err)
	}
	return vote, nil
}

func (s *Service) GetVotes() ([]domain.Vote, error) {
	votes, err := s.repo.GetVotes()
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
	return s.repo.AddVote(v)
}

// AddVotes publishes the votes to the queue
func (s *Service) AddVotes(v []domain.Vote) error {
	return s.repo.AddVotes(v)
}
