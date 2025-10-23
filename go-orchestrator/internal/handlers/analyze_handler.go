package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"autopilot-engineer/go-orchestrator/internal/core"
	"autopilot-engineer/go-orchestrator/internal/services"
)

type AnalyzeHandler struct {
	LangGraph services.LangGraphService
}

func NewAnalyzeHandler(langGraph services.LangGraphService) *AnalyzeHandler {
	return &AnalyzeHandler{LangGraph: langGraph}
}

func (h *AnalyzeHandler) HandleAnalyze(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	// Use pure function (no HTTP inside it)
	prepared := core.ProcessInput(payload)

	resp, err := h.LangGraph.SendToLangGraph(context.Background(), prepared)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   resp,
	})
}
