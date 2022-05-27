package repository

import "github.com/guibedin/poll/web/domain"

type Repository interface {
	domain.PollRepository
	domain.OptionRepository
	domain.VoteRepository
}
