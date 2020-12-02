package main

import (
	"github.com/gofiber/fiber/v2"
)

type Ctx = *fiber.Ctx

func main() {
	app := fiber.New()

	app.Get("/", func(c Ctx) error {
		return c.SendString("Hello World!")
	})
	app.Listen(":3000")
}
