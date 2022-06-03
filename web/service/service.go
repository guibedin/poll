package service

import "github.com/guibedin/poll/web/repository"

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return Service{repo}
}
