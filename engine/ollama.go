package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var _ Engine = (*ollamaEngine)(nil)

type ollamaEngine struct {
	Host  string
	Model string
}

func NewOllama(host, model string) ollamaEngine {
	return ollamaEngine{
		Host:  host + "/api/generate",
		Model: model,
	}
}

func (ollama *ollamaEngine) GetCommitMessage(gitDiff, style string) (string, error) {
	return ollama.request(gitDiff, style)
}

func (ollama *ollamaEngine) request(diff, style string) (string, error) {
	payload := map[string]any{
		"model":  ollama.Model,
		"prompt": createPrompt(diff, style),
		"format": "json",
		"stream": false,
		"options": map[string]any{
			"temperature": 0,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(ollama.Host, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}

	if res == nil || res["response"] == nil {
		return "", fmt.Errorf("ollama %s", res["error"])
	}

	// `response` is a JSON string
	// https://github.com/ollama/ollama/blob/main/docs/api.md#response-2
	var message struct {
		CommitMessage string `json:"commit_message"`
	}
	if err := json.Unmarshal([]byte(res["response"].(string)), &message); err != nil {
		return "", err
	}

	return message.CommitMessage, nil
}
