package read

type Service interface {
	GetPoll(string) (Poll, error)
	GetAllPolls() []Poll
}
