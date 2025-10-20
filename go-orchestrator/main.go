package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": "go-orchestrator"})
	})

	log.Println(" Go Orchestrator running on :8080")
	app.Listen(":8080")
}
