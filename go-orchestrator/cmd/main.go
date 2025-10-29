package main

import (
	"log"

	"autopilot-engineer/go-orchestrator/internal/handlers"
	"autopilot-engineer/go-orchestrator/internal/routes"
	"autopilot-engineer/go-orchestrator/internal/services"
	"autopilot-engineer/go-orchestrator/internal/db"
)

func main() {
	langGraphService := services.NewLangGraphService()
	analyzeHandler := handlers.NewAnalyzeHandler(langGraphService)

	router := routes.SetupRouter(analyzeHandler)
	db.InitMongo()
	log.Println("🚀 Starting Go orchestrator on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
