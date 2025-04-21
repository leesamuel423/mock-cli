package models_test

import (
	"testing"

	"interview-cli/internal/models"
)

func TestIsValidCategory(t *testing.T) {
	tests := []struct {
		name     string
		category string
		want     bool
	}{
		{"Valid Behavioral", "behavioural", true},
		{"Valid Technical", "technical", true},
		{"Invalid Category", "other", false},
		{"Empty Category", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := models.IsValidCategory(tt.category); got != tt.want {
				t.Errorf("IsValidCategory(%q) = %v, want %v", tt.category, got, tt.want)
			}
		})
	}
}

func TestHasTag(t *testing.T) {
	question := models.Question{
		Text:     "Test question",
		Category: models.CategoryTechnical,
		Tags:     []string{"golang", "testing", "interview"},
	}

	tests := []struct {
		name string
		tag  string
		want bool
	}{
		{"Has Tag", "golang", true},
		{"Has Another Tag", "testing", true},
		{"Does Not Have Tag", "database", false},
		{"Empty Tag", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := question.HasTag(tt.tag); got != tt.want {
				t.Errorf("Question.HasTag(%q) = %v, want %v", tt.tag, got, tt.want)
			}
		})
	}
}

func TestHasAllTags(t *testing.T) {
	question := models.Question{
		Text:     "Test question",
		Category: models.CategoryTechnical,
		Tags:     []string{"golang", "testing", "interview"},
	}

	tests := []struct {
		name string
		tags []string
		want bool
	}{
		{"No Tags", []string{}, true},
		{"Single Matching Tag", []string{"golang"}, true},
		{"Multiple Matching Tags", []string{"golang", "testing"}, true},
		{"All Tags", []string{"golang", "testing", "interview"}, true},
		{"One Non-matching Tag", []string{"golang", "database"}, false},
		{"All Non-matching Tags", []string{"database", "frontend"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := question.HasAllTags(tt.tags); got != tt.want {
				t.Errorf("Question.HasAllTags(%v) = %v, want %v", tt.tags, got, tt.want)
			}
		})
	}
}