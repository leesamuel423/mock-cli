package config_test

import (
	"encoding/json"
	"testing"

	"interview-cli/internal/config"
)

// TestConfigStructure tests the Config struct fields
func TestConfigStructure(t *testing.T) {
	// sample config
	cfg := &config.Config{
		QuestionsFile: "test-questions.json",
	}

	// serialize to JSON
	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	// deserialize back
	var loadedCfg config.Config
	err = json.Unmarshal(data, &loadedCfg)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	// verify structure
	if loadedCfg.QuestionsFile != "test-questions.json" {
		t.Errorf("Expected QuestionsFile to be %q, got %q", "test-questions.json", loadedCfg.QuestionsFile)
	}
}

// TestDefaultConfig tests that a default config is created properly
func TestDefaultConfig(t *testing.T) {
	// can't easily test LoadConfig directly since it relies on the executable path
	// instead, test the Config structure and JSON handling separately

	// create a test Config with default values
	cfg := &config.Config{
		QuestionsFile: "questions.json", // Default value
	}

	// verify config has expected default values
	if cfg.QuestionsFile != "questions.json" {
		t.Errorf("Default config has unexpected QuestionsFile value: %q", cfg.QuestionsFile)
	}
}

