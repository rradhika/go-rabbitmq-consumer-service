package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rradhika/go-rabbitmq-consumer-service/config"
	"github.com/rradhika/go-rabbitmq-consumer-service/routes"
	"github.com/rradhika/go-rabbitmq-consumer-service/utils"
)

func main() {
	utils.LoadEnv()
	utils.InitDB()
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
