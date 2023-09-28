package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World")
	// })

	app.Static("/", "./public")

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("Value : " + c.Params("value"))
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg)
	})

	app.Listen(":3000")
}
