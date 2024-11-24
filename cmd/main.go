package main

import (
	database "main/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hey bud fuk!")
	})

	app.Listen(":3000")
}
