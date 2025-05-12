package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rradhika/go-rabbitmq-producer/config"
	"github.com/rradhika/go-rabbitmq-producer/routes"
	"github.com/rradhika/go-rabbitmq-producer/utils"
)

func main() {
	utils.LoadEnv()
	utils.InitRabbitMQ()
	utils.ConsumeQueue()
	conf := config.GetConfig()

	app := fiber.New()

	routes.Register(app)

	fmt.Printf("🚀 %s running on http://localhost:%s\n", conf.AppName, conf.Port)
	err := app.Listen(":" + conf.Port)
	if err != nil {
		panic(err)
	}
}
