package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendString("v1.0.0")
	})

	app.Listen(":3000")
}
