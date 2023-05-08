package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// cfg := config.LoadConfig()

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong!")
	})

	app.Listen(":3000")
}
