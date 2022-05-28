package service

import "github.com/guibedin/poll/web/domain"

func (s *Service) GetPoll(id int) (domain.Poll, error) {
	poll, err := s.repo.GetPoll(id)
	if err != nil {
		panic(err)
	}
	return poll, nil
}

// GetPolls returns all Polls from the database
func (s *Service) GetPolls() ([]domain.Poll, error) {
	polls, err := s.repo.GetPolls()
	if err != nil {
		panic(err)
	}
	return polls, nil
}

// AddPoll adds the new poll to the database
func (s *Service) AddPoll(p domain.Poll) int {
	return s.repo.AddPoll(p)
}
