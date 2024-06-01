package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
)

type Config struct {
	OllamaHost  string `json:"ollama_host" env:"OLLAMA_HOST"`
	OllamaModel string `json:"ollama_model" env:"OLLAMA_MODEL"`
}

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
			}
		}
	}

	return nil
}

func (cfg *Config) Save() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644) //os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
