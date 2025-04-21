package repository_test

import (
	"encoding/json"
	"os"
	"testing"

	"interview-cli/internal/models"
	"interview-cli/internal/repository"
)

// helper function to create a temporary file with test questions
func createTempQuestionFile(t *testing.T, questions []models.Question) string {
	t.Helper()

	// create temp file
	tmpFile, err := os.CreateTemp("", "questions-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// close the file but keep the path
	filename := tmpFile.Name()
	tmpFile.Close()

	// if questions are provided, write them to the file
	if len(questions) > 0 {
		data, err := json.MarshalIndent(questions, "", " ")
		if err != nil {
			t.Fatalf("Failed to marshal questions: %v", err)
		}

		if err := os.WriteFile(filename, data, 0644); err != nil {
			t.Fatalf("Failed to write questions to temp file: %v", err)
		}
	}

	// register cleanup function
	t.Cleanup(func() {
		os.Remove(filename)
	})

	return filename
}

func TestLoadQuestions(t *testing.T) {
	// test data
	testQuestions := []models.Question{
		{Text: "Q1", Category: models.CategoryTechnical, Tags: []string{"golang"}},
		{Text: "Q2", Category: models.CategoryBehavioural, Tags: []string{"leadership"}},
	}

	filePath := createTempQuestionFile(t, testQuestions)
	repo := repository.NewQuestionRepository(filePath)

	// test loading questions
	questions, err := repo.LoadQuestions()
	if err != nil {
		t.Fatalf("Failed to load questions: %v", err)
	}

	// verify questions
	if len(questions) != len(testQuestions) {
		t.Errorf("Expected %d questions, got %d", len(testQuestions), len(questions))
	}

	for i, q := range questions {
		if q.Text != testQuestions[i].Text {
			t.Errorf("Question %d text mismatch: expected %q, got %q", i, testQuestions[i].Text, q.Text)
		}
		if q.Category != testQuestions[i].Category {
			t.Errorf("Question %d category mismatch: expected %q, got %q", i, testQuestions[i].Category, q.Category)
		}
	}
}

func TestSaveQuestions(t *testing.T) {
	filePath := createTempQuestionFile(t, nil)
	repo := repository.NewQuestionRepository(filePath)

	// test data to save
	testQuestions := []models.Question{
		{Text: "Save Q1", Category: models.CategoryTechnical, Tags: []string{"golang"}},
		{Text: "Save Q2", Category: models.CategoryBehavioural, Tags: []string{"teamwork"}},
	}

	// save questions
	err := repo.SaveQuestions(testQuestions)
	if err != nil {
		t.Fatalf("Failed to save questions: %v", err)
	}

	// verify by loading them back
	loadedQuestions, err := repo.LoadQuestions()
	if err != nil {
		t.Fatalf("Failed to load questions after saving: %v", err)
	}

	// verify content
	if len(loadedQuestions) != len(testQuestions) {
		t.Errorf("Expected %d questions, got %d", len(testQuestions), len(loadedQuestions))
	}

	for i, q := range loadedQuestions {
		if q.Text != testQuestions[i].Text {
			t.Errorf("Saved question %d text mismatch: expected %q, got %q", i, testQuestions[i].Text, q.Text)
		}
	}
}

func TestFindQuestionsByFilter(t *testing.T) {
	// Test data
	testQuestions := []models.Question{
		{Text: "Tech Q1", Category: models.CategoryTechnical, Tags: []string{"golang", "testing"}},
		{Text: "Tech Q2", Category: models.CategoryTechnical, Tags: []string{"python", "algorithms"}},
		{Text: "Behavioral Q1", Category: models.CategoryBehavioural, Tags: []string{"leadership", "teamwork"}},
		{Text: "Behavioral Q2", Category: models.CategoryBehavioural, Tags: []string{"conflict", "teamwork"}},
	}

	filePath := createTempQuestionFile(t, testQuestions)
	repo := repository.NewQuestionRepository(filePath)

	// test cases
	tests := []struct {
		name     string
		category string
		tags     []string
		want     int
	}{
		{"All Questions", "", []string{}, 4},
		{"Technical Category", models.CategoryTechnical, []string{}, 2},
		{"Behavioral Category", models.CategoryBehavioural, []string{}, 2},
		{"Single Tag Filter", "", []string{"teamwork"}, 2},
		{"Category and Tag Filter", models.CategoryBehavioural, []string{"teamwork"}, 2},
		{"Multiple Tags Filter", "", []string{"golang", "testing"}, 1},
		{"No Matching Questions", "", []string{"nonexistent"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			questions, err := repo.FindQuestionsByFilter(tt.category, tt.tags)
			if err != nil {
				t.Fatalf("Failed to find questions by filter: %v", err)
			}

			if len(questions) != tt.want {
				t.Errorf("FindQuestionsByFilter(%q, %v) returned %d questions, want %d",
					tt.category, tt.tags, len(questions), tt.want)
			}
		})
	}
}

