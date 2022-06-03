package service

import "github.com/guibedin/poll/web/domain"

// GetOptionsByPollID returns all options from a given Poll
func (s *Service) GetOptionsByPollID(id int) ([]domain.Option, error) {
	options, err := s.repo.GetOptionsByPollID(id)
	if err != nil {
		panic(err)
	}
	return options, nil
}

// AddOption adds an option to the database
func (s *Service) AddOption(o domain.Option, pollId int) int {
	return s.repo.AddOption(o, pollId)
}
