package services

import (
	"fmt"
)

type AnalyzeService struct{}

func NewAnalyzeService() *AnalyzeService {
	return &AnalyzeService{}
}

func (s *AnalyzeService) AnalyzeText(input string) (string, error) {
	// TODO: Call LangGraph engine later
	fmt.Println("Received text:", input)
	return fmt.Sprintf("Mock analysis for: %s", input), nil
}
