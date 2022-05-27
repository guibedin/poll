package domain

import (
	"time"
)

type Poll struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	IsActive         bool      `json:"is_active"`
	IsMultipleChoice bool      `json:"is_multiple_choice"`
	Options          []Option  `json:"options"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type PollRepository interface {
	GetPoll(ID int) (Poll, error)
	GetPolls() ([]Poll, error)
	AddPoll(p Poll) int
}
