package config

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var configKey = "kemit-config"

type Config struct {
	OllamaHost  string `json:"ollama_host"`
	OllamaModel string `json:"ollama_model"`
}

func (cfg *Config) SaveConfig() error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "darwin":
		return saveConfigMacOS(string(jsonData), configKey)
	case "linux":
		return saveConfigLinux(string(jsonData), configKey)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func (cfg *Config) LoadConfig() error {
	var output []byte
	var err error
	switch runtime.GOOS {
	case "darwin":
		output, err = loadConfigMacOS(configKey)
	case "linux":
		output, err = loadConfigLinux(configKey)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal(output, cfg)
	if err != nil {
		return err
	}

	return nil
}
