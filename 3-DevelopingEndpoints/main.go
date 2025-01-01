package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()

	app.Get("/", adaptor.HTTPHandler(templ.Handler(OutOfBand())))

	app.Get("/demo", adaptor.HTTPHandler(templ.Handler(Demo())))

	app.Listen(":3000")
}
