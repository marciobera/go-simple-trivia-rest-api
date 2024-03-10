package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marciobera/go-simple-trivia-rest-api/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World, again!")
	})

	app.Listen(":3000")
}