package domain

type Vote struct {
	ID       int    `json:"id"`
	OptionId int    `json:"option_id"`
	PollId   int    `json:"poll_id"`
	Voter    string `json:"voter"`
}
