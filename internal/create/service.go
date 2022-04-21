package create

import (
	"github.com/guibedin/voting/internal/read"
)

type Service interface {
	CreatePoll(Poll) error
}

type Repository interface {
	CreatePoll(Poll) error
	GetAllPolls() []read.Poll
}
