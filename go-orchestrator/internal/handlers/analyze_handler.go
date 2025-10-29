package handlers

import (
	"context"
	"net/http"
	"os"

	"autopilot-engineer/go-orchestrator/internal/core"
	"autopilot-engineer/go-orchestrator/internal/services"
	"autopilot-engineer/go-orchestrator/internal/utils"

	"github.com/gin-gonic/gin"
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
	// --- NEW: handle repo URL ---
	if repoURL, ok := payload["repo_url"].(string); ok && repoURL != "" {
		repoPath, err := utils.CloneRepo(repoURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clone repo", "details": err.Error()})
			return
		}
		defer func() { _ = os.RemoveAll(repoPath) }()

		summary := utils.SummarizeRepo(repoPath)
		prepared := map[string]interface{}{
			"repo_summary": summary,
		}

		resp, err := h.LangGraph.SendToLangGraph(context.Background(), prepared)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok", "data": resp})
		services.SaveRepoSummary(repoURL, resp)


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
