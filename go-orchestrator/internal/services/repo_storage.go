package services

import (
	"context"
	"log"
	"time"

	"autopilot-engineer/go-orchestrator/internal/db"
)

type RepoSummary struct {
	RepoURL     string                 `bson:"repo_url"`
	AnalyzedAt  time.Time              `bson:"analyzed_at"`
	SummaryData map[string]interface{} `bson:"summary_data"`
}

func SaveRepoSummary(repoURL string, summary map[string]interface{}) error {
	collection := db.GetCollection("autopilot", "repo_summaries")

	doc := RepoSummary{
		RepoURL:     repoURL,
		AnalyzedAt:  time.Now(),
		SummaryData: summary,
	}

	_, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Println("❌ Failed to save repo summary:", err)
		return err
	}

	log.Println("✅ Repo summary saved to MongoDB:", repoURL)
	return nil
}
