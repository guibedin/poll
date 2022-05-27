package consumer

import (
	"encoding/json"
	"log"

	"github.com/guibedin/poll/consumer/domain"
	"github.com/guibedin/poll/consumer/repository"
	"github.com/streadway/amqp"
)

type Consumer struct {
	repo repository.Repository
	mq   *amqp.Connection
}

func (c *Consumer) SetRepository(repo repository.Repository) {
	c.repo = repo
}

func (c *Consumer) SetMQ(mq *amqp.Connection) {
	c.mq = mq
}

func (c *Consumer) Receive() {

	ch, err := c.mq.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"votes", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	// Create a go routine that scans through messages
	// The `for` blocks when trying to read from channel `msgs`, waiting for a new message
	go func() {
		for d := range msgs {
			var vote domain.Vote
			json.Unmarshal(d.Body, &vote)
			log.Printf("Received a message: %+v", vote)
			c.repo.AddVote(vote)
			log.Println("Vote processed!")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// Blocks execution of the code, because nothing is sent to the channel forever.
	// This is done to keep the go routine that was created above running.
	<-forever
}

func New() *Consumer {
	c := &Consumer{}
	return c
}
