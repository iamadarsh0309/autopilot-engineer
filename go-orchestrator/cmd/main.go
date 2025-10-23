package main

import (
	"log"

	"autopilot-engineer/go-orchestrator/internal/handlers"
	"autopilot-engineer/go-orchestrator/internal/routes"
	"autopilot-engineer/go-orchestrator/internal/services"
)

func main() {
	langGraphService := services.NewLangGraphService()
	analyzeHandler := handlers.NewAnalyzeHandler(langGraphService)

	router := routes.SetupRouter(analyzeHandler)

	log.Println("🚀 Starting Go orchestrator on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
