package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

// type option struct {
// 	Description string
// 	Count       int64
// }

// type poll struct {
// 	Title   string
// 	Options []option
// }

type vote struct {
	Option string `json:"option"`
}

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

func votes(w http.ResponseWriter, r *http.Request) {
	// Decodes body to struct
	decoder := json.NewDecoder(r.Body)
	var v vote
	err := decoder.Decode(&v)
	failOnError(err, "Failed to decode request body.")

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

	// Publish message
	body := v.Option
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	// Write response
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Thank you for your vote!"
	json.NewEncoder(w).Encode(resp)
}

func results(w http.ResponseWriter, r *http.Request) {
	// Write response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["results"] = "Option 1: 1; Option 2: 2"
	json.NewEncoder(w).Encode(resp)
}

func main() {

	http.HandleFunc("/api/votes", votes)
	http.HandleFunc("/api/results", results)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
