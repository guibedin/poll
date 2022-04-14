package main

import "github.com/streadway/amqp"

func connectToAmqp() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ.")
	return conn
}

func publish(body []byte) error {
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
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	return nil
}
