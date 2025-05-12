// handlers/rabbit.go
package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rradhika/go-rabbitmq-producer/utils"
)

func SendMessage(c *fiber.Ctx) error {
	msg := c.Params("msg")
	totalStr := c.Params("total")
	total, err := strconv.Atoi(totalStr)
	if err != nil {
		return c.Status(400).SendString("Invalid total value")
	}

	for i := 0; i < total; i++ {
		err := utils.Channel.Publish(
			"", utils.Queue.Name, false, false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
				Timestamp:   time.Now(),
			},
		)

		if err != nil {
			return c.Status(500).SendString("Failed to publish message")
		}
	}

	return c.SendString("Message sent to RabbitMQ: " + msg)
}
