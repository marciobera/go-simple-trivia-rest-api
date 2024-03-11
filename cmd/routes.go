package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marciobera/go-simple-trivia-rest-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/facts", handlers.ListFacts)
	app.Post("/facts", handlers.CreateFact)
}