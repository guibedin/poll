package repository

import "github.com/guibedin/poll/consumer/domain"

type Repository interface {
	AddVote(v domain.Vote) error
}
