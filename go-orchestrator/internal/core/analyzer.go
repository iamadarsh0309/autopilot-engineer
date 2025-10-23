package core

func ProcessInput(payload map[string]interface{}) map[string]interface{} {
	query, _ := payload["query"].(string)
	return map[string]interface{}{
		"text": query,
	}
}
