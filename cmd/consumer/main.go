package main

import (
	"github.com/guibedin/poll/consumer/consumer"
	"github.com/guibedin/poll/consumer/mq"
	"github.com/guibedin/poll/consumer/repository"
)

<<<<<<< HEAD
=======
type storageType int

const (
	sql storageType = iota
	file
)

>>>>>>> 0bfecc2a6d667f77f231ab1746dc1b360f89c8b6
func main() {
	consumer := consumer.New()

	consumer.SetMQ(mq.NewMQConnection())
	// Set repository type
	consumer.SetRepository(repository.New(repository.Sql))

	consumer.Receive("votes")
}
