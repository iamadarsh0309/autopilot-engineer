package main

import (
	"autopilot-engineer/go-orchestrator/config"
	"autopilot-engineer/go-orchestrator/internal/handlers"
	"autopilot-engineer/go-orchestrator/internal/middleware"
	"autopilot-engineer/go-orchestrator/internal/services"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()
	app := fiber.New()

	// Middleware
	app.Use(middleware.RequestLogger())

	// Routes
	analyzeService := services.NewAnalyzeService()
	analyzeHandler := handlers.NewAnalyzeHandler(analyzeService)

	app.Post("/api/analyze", analyzeHandler.Analyze)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"service": "go-orchestrator", "status": "ok"})
	})

	// Graceful shutdown
	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()
}
