package mq

import (
	"log"

	"github.com/streadway/amqp"
)

func NewMQConnection() *amqp.Connection {
	// Get MQ connection
	mq, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	log.Println("Consumer connected to MQ")

	return mq
}
