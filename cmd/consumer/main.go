package main

import (
	"github.com/guibedin/poll/consumer/consumer"
	"github.com/guibedin/poll/consumer/mq"
	"github.com/guibedin/poll/consumer/repository"
)

type storageType int

const (
	sql  storageType = 1
	file storageType = 2
)

func main() {
	consumer := consumer.New()

	var repo repository.Repository
	repoType := sql
	switch repoType {
	case sql:
		repo = repository.NewSqlRepository()
	case file:
		repo = repository.NewFileRepository()
	}
	consumer.SetMQ(mq.NewMQConnection())
	consumer.SetRepository(repo)

	consumer.Receive()
}
