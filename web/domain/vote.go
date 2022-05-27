package domain

type Vote struct {
	ID       int    `json:"id"`
	OptionId int    `json:"option_id"`
	PollId   int    `json:"poll_id"`
	Voter    string `json:"voter"`
}

type VoteRepository interface {
	GetVote() (Vote, error)
	GetVotes() ([]Vote, error)
	GetVotesByOptionID(id int) ([]Vote, error)
	GetVoteCountByOptionID(id int) (int, error)
	AddVote(v Vote) error
	AddVotes(v []Vote) error
}
