package domain

import "time"

type Option struct {
	ID        int       `json:"id"`
	PollId    int       `json:"poll_id,omitempty"`
	Title     string    `json:"title"`
	Votes     int       `json:"votes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OptionRepository interface {
	GetOption() (Option, error)
	GetOptionsByPollID(id int) ([]Option, error)
	AddOption(o Option, pollId int) int
}
