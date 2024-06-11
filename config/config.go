package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type Config struct {
	Provider string `json:"provider" env:"PROVIDER" default:"ollama"`

	OllamaHost  string `json:"ollama_host" env:"OLLAMA_HOST"`
	OllamaModel string `json:"ollama_model" env:"OLLAMA_MODEL"`
}

var (
	ollamaMissingDataError   = errors.New("missing config. ollama_model or ollama_host")
	unsupportedProviderError = errors.New("unsupported provider")
)

func (cfg *Config) Load() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(configPath)
	if err == nil {
		err := json.Unmarshal(file, &cfg)
		if err != nil {
			return err
		}
	}

	// Override with environment variables
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envTag := field.Tag.Get("env")
		if envTag != "" {
			envValue := os.Getenv(envTag)
			if envValue != "" {
				v.Field(i).SetString(envValue)
				continue
			}
		}

		defaultTag := field.Tag.Get("default")
		if v.Field(i).String() == "" && defaultTag != "" {
			v.Field(i).SetString(defaultTag)
		}
	}

	return nil
}

func (cfg *Config) Save() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) Validate() error {
	switch strings.ToLower(cfg.Provider) {
	case "ollama":
		if cfg.OllamaHost == "" || cfg.OllamaModel == "" {
			return ollamaMissingDataError
		}
	default:
		return unsupportedProviderError
	}

	return nil
}

func getConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(configDir, "kemit", "conf")

	_, err = os.Stat(configPath)
	if err != nil {
		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return "", err
		}

		file, err := os.Create(configPath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		_, err = file.Write([]byte("{}"))
		if err != nil {
			return "", err
		}
	}

	return configPath, nil
}
