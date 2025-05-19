// utils/rabbitmq.go
package utils

import (
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp091.Channel
var Queue amqp091.Queue

func InitRabbitMQ() {
	conn, err := amqp091.Dial(os.Getenv("AMQP_DSN"))
	if err != nil {
		log.Fatalf("RabbitMQ connection error: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Channel open error: %s", err)
	}

	q, err := ch.QueueDeclare(
		"testing_queue", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Queue declare error: %s", err)
	}

	RabbitChannel = ch
	Queue = q
}

func ConsumeQueue() {
	msgs, err := RabbitChannel.Consume(
		Queue.Name,
		"",
		true, // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Failed to register consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			content := string(msg.Body)

			// Insert into DB
			res, err := DB.Exec(`INSERT INTO public.messages (content, created_at) VALUES ($1, $2)`, content, time.Now())
			if err != nil {
				log.Printf("❌ DB insert error: %v", err)
			} else {
				n, _ := res.RowsAffected()
				log.Printf("✅ Inserted %d row(s): %s", n, content)
			}

			var count int
			_ = DB.Get(&count, "SELECT COUNT(*) FROM public.messages")
			log.Println("📊 Total rows in DB after insert:", count)
		}
	}()

	log.Println("📡 RabbitMQ consumer started")
}
