package poll

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Option struct {
	pollId      string
	Description string `json:"description"`
	Count       int64  `json:"count"`
}

type Poll struct {
	Title     string    `json:"title"`
	Options   []Option  `json:"options"`
	CreatedAt time.Time `json:"createdAt"`
	IsActive  bool      `json:"isActive"`
}

func connectToAmqp() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ.")
	return conn
}

func New(title string, options []Option) *Poll {
	return &Poll{
		Title:     title,
		Options:   options,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
}

func (p *Poll) Vote(option string) bool {
	// Connecto to RabbitMQ
	conn := connectToAmqp()
	defer conn.Close()

	// Open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel.")
	defer ch.Close()

	// Declare Queue
	q, err := ch.QueueDeclare(
		"votes", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(option),
		})
	failOnError(err, "Failed to publish a message")

	return true
}
