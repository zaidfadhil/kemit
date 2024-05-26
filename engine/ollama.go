package engine

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var _ Engine = (*ollamaEngine)(nil)

type ollamaEngine struct {
	Host  string
	Model string
	Diff  string
}

func NewOllama(diff string) *ollamaEngine {
	return &ollamaEngine{
		Host:  "http://192.168.0.107:11435/api/generate",
		Model: "llama3",
		Diff:  diff,
	}
}

func (ollama *ollamaEngine) GetCommit() (string, error) {
	return ollama.request()
}

func (ollama *ollamaEngine) request() (string, error) {

	payload := map[string]any{
		"model":  ollama.Model,
		"prompt": createPrompt(ollama.Diff),
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
	json.NewDecoder(resp.Body).Decode(&res)

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
