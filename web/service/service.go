package service

import "github.com/guibedin/poll/web/repository"

type Service struct {
	repo repository.Repository
}

func NewSqlService() Service {
	return Service{repository.NewSqlRepository()}
}

func NewFileService() Service {
	return Service{repository.NewFileRepository()}
}
