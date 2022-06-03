package main

import (
	"github.com/guibedin/poll/consumer/consumer"
	"github.com/guibedin/poll/consumer/mq"
	"github.com/guibedin/poll/consumer/repository"
)

type storageType int

const (
	sql storageType = iota
	file
)

func main() {
	consumer := consumer.New()

	consumer.SetMQ(mq.NewMQConnection())
	// Set repository type
	consumer.SetRepository(repository.New(repository.Sql))

	consumer.Receive("votes")
}
