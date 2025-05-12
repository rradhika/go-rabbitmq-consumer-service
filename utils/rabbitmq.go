// utils/rabbitmq.go
package utils

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var Channel *amqp091.Channel
var Queue amqp091.Queue

func InitRabbitMQ() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("RabbitMQ connection error: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Channel open error: %s", err)
	}

	q, err := ch.QueueDeclare(
		"hello_queue", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Queue declare error: %s", err)
	}

	Channel = ch
	Queue = q
}
