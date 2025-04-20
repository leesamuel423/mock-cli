package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// application configuration ... stores location of the questions file
type Config struct {
	QuestionsFile string `json:"questionsFile"` // path to file
}

// load config from disk; if no config, create default config
// returns pointer to config and error
func LoadConfig() (*Config, error) {
	// Get the executable directory
	execPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to get executable path: %w", err)
	}
	execDir := filepath.Dir(execPath)

	// Use executable directory for config
	configPath := filepath.Join(execDir, "config.json")

	// init default config structure
	config := &Config{
		QuestionsFile: filepath.Join(execDir, "questions.json"),
	}

	// check if config file exists
	if _, err := os.Stat(configPath); err == nil {
		// read and parse JSON
		data, err := os.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse config JSON: %w", err)
		}

		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("failed to parse config JSON: %w", err)
		}
	} else { // if config file doesn't exist, create w/ default values
		data, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal config to JSON: %w", err)
		}

		// write default config to disk
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to write config to file: %w", err)
		}
	}

	// check if questions file exists
	if _, err := os.Stat(config.QuestionsFile); os.IsNotExist(err) {
		emptyQuestions := []any{}

		// convert empty slice to formatted JSON
		data, err := json.MarshalIndent(emptyQuestions, "", " ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal empty questions to JSON: %w", err)
		}

		// write empty question array to disk
		if err := os.WriteFile(config.QuestionsFile, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to create questions file: %w", err)
		}
	}

	// return loaded or created config
	return config, nil
}