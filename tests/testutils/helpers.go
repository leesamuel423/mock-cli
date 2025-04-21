package testutils

import (
	"encoding/json"
	"os"
	"testing"

	"interview-cli/internal/models"
)

// CreateTempFile creates a temporary file and returns its path
// The file will be automatically cleaned up when the test completes
func CreateTempFile(t *testing.T, prefix string) string {
	t.Helper()
	
	tmpFile, err := os.CreateTemp("", prefix)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	
	filename := tmpFile.Name()
	tmpFile.Close()
	
	t.Cleanup(func() {
		os.Remove(filename)
	})
	
	return filename
}

// WriteQuestionsToFile writes a slice of questions to a JSON file
func WriteQuestionsToFile(t *testing.T, filename string, questions []models.Question) {
	t.Helper()
	
	data, err := json.MarshalIndent(questions, "", " ")
	if err != nil {
		t.Fatalf("Failed to marshal questions: %v", err)
	}
	
	if err := os.WriteFile(filename, data, 0644); err != nil {
		t.Fatalf("Failed to write questions to file: %v", err)
	}
}

// ReadQuestionsFromFile reads a slice of questions from a JSON file
func ReadQuestionsFromFile(t *testing.T, filename string) []models.Question {
	t.Helper()
	
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read questions file: %v", err)
	}
	
	var questions []models.Question
	if err := json.Unmarshal(data, &questions); err != nil {
		t.Fatalf("Failed to parse questions JSON: %v", err)
	}
	
	return questions
}