package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LangGraphService interface {
	SendToLangGraph(ctx context.Context, payload interface{}) (map[string]interface{}, error)
}

type langGraphService struct {
	baseURL string
	client  *http.Client
}

func NewLangGraphService() LangGraphService {
	return &langGraphService{
		baseURL: "http://langgraph-engine:8000", // 👈 FIXED port
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *langGraphService) SendToLangGraph(ctx context.Context, payload interface{}) (map[string]interface{}, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/analyze", s.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling LangGraph engine: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("LangGraph returned status: %v", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return result, nil
}
