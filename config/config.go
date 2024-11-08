package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type Config struct {
	Provider string `json:"provider" env:"PROVIDER" default:"ollama"`

	OllamaHost  string `json:"ollama_host" env:"OLLAMA_HOST"`
	OllamaModel string `json:"ollama_model" env:"OLLAMA_MODEL"`

	CommitStyle string `json:"commit_style" env:"COMMIT_STYLE" default:"conventional-commit"`
}

var (
	ErrMissingOllamaData   = errors.New("missing config. ollama_model or ollama_host")
	ErrUnsupportedProvider = errors.New("unsupported provider")
)

func (cfg *Config) Load() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(filepath.Clean(configPath))
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
				setFieldValue(v.Field(i), envValue)
				continue
			}
		}

		defaultTag := field.Tag.Get("default")
		if v.Field(i).IsZero() && defaultTag != "" {
			setFieldValue(v.Field(i), defaultTag)
		}
	}

	return nil
}

func (cfg *Config) Save() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath.Clean(configPath), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
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
			return ErrMissingOllamaData
		}
	default:
		return ErrUnsupportedProvider
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
		err = os.MkdirAll(filepath.Dir(configPath), 0750)
		if err != nil {
			return "", err
		}

		file, err := os.OpenFile(filepath.Clean(configPath), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
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

func setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			field.SetInt(intValue)
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			field.SetBool(boolValue)
		}
	}
}
