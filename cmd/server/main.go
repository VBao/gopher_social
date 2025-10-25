package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "okla",
		})
	})

	error := app.Listen(":8099")

	if error != nil {
		log.Fatal(error)
	}
}
