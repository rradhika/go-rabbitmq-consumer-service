package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rradhika/go-rabbitmq-producer/handlers"
)

func Register(app *fiber.App) {
	app.Get("/", handlers.Hello)
	app.Get("/send/:msg/:total", handlers.SendMessage)
}
