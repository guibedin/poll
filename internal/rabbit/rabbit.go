package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectToAmqp() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ.")
	return conn
}

func Publish(body []byte) error {
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

// func consume() {
// 	conn := connectToAmqp()
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	failOnError(err, "failed to open a channel")
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"votes", // name
// 		false,   // durable
// 		false,   // delete when unused
// 		false,   // exclusive
// 		false,   // no-wait
// 		nil,     // arguments
// 	)
// 	failOnError(err, "Failed to declare a queue")

// 	forever := make(chan bool)

// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	failOnError(err, "Failed to register a consumer")

// 	go func() {
// 		for d := range msgs {
// 			var v vote
// 			err := json.Unmarshal(d.Body, &v)
// 			failOnError(err, "failed to unmarshall body")
// 			log.Printf("Received a message: %s", v)
// 			countVote(v)
// 		}
// 	}()

// 	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
// 	<-forever
// }
