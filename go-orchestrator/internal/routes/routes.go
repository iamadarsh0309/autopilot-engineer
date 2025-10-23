package routes

import (
	"github.com/gin-gonic/gin"
	"autopilot-engineer/go-orchestrator/internal/handlers"
)

func SetupRouter(analyzeHandler *handlers.AnalyzeHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.POST("/analyze", analyzeHandler.HandleAnalyze)
	return router
}
