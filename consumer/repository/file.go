package repository

import "github.com/guibedin/poll/consumer/domain"

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (f *fileRepository) AddVote(v domain.Vote) error {
	return nil
}
