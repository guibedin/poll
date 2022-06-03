package repository

import "github.com/guibedin/poll/consumer/domain"

type RepositoryType int

const (
	Sql  RepositoryType = iota // = 1
	File                       // = 2
)

type Repository interface {
	AddVote(v domain.Vote) error
}

// Sql creates the sqlRepository, which implements Repostitory, that will be used in the server
func New(repoType RepositoryType) Repository {

	switch repoType {
	case Sql:
		return NewSqlRepository()
	case File:
		return NewFileRepository()
	default:
		return NewFileRepository()
	}
}
