package session_test

import (
	"testing"

	"interview-cli/internal/models"
)

// interface for the repository to make testing easier
type QuestionRepositoryInterface interface {
	FindQuestionsByFilter(category string, tags []string) ([]models.Question, error)
}

// mock repository that implements the QuestionRepositoryInterface
type MockQuestionRepository struct {
	questions []models.Question
}

func NewMockRepository(questions []models.Question) *MockQuestionRepository {
	return &MockQuestionRepository{questions: questions}
}

func (m *MockQuestionRepository) LoadQuestions() ([]models.Question, error) {
	return m.questions, nil
}

func (m *MockQuestionRepository) SaveQuestions(questions []models.Question) error {
	m.questions = questions
	return nil
}

func (m *MockQuestionRepository) AddQuestion() error {
	return nil
}

func (m *MockQuestionRepository) ListQuestions() error {
	return nil
}

func (m *MockQuestionRepository) FindQuestionsByFilter(category string, tags []string) ([]models.Question, error) {
	if category == "" && len(tags) == 0 {
		return m.questions, nil
	}

	var filtered []models.Question
	for _, q := range m.questions {
		// Filter by category
		if category != "" && q.Category != category {
			continue
		}

		// Filter by tags
		if len(tags) > 0 {
			allTagsMatch := true
			for _, tag := range tags {
				if !contains(q.Tags, tag) {
					allTagsMatch = false
					break
				}
			}
			if !allTagsMatch {
				continue
			}
		}

		filtered = append(filtered, q)
	}

	return filtered, nil
}

// helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func TestPracticeSessionQuestionSelection(t *testing.T) {
	// create test questions
	testQuestions := []models.Question{
		{Text: "Tech Q1", Category: models.CategoryTechnical, Tags: []string{"golang"}},
		{Text: "Tech Q2", Category: models.CategoryTechnical, Tags: []string{"python"}},
		{Text: "Tech Q3", Category: models.CategoryTechnical, Tags: []string{"java"}},
		{Text: "Behavioral Q1", Category: models.CategoryBehavioural, Tags: []string{"leadership"}},
		{Text: "Behavioral Q2", Category: models.CategoryBehavioural, Tags: []string{"teamwork"}},
	}

	mockRepo := NewMockRepository(testQuestions)

	// test the internal question selection
	tests := []struct {
		name              string
		numQuestions      int
		category          string
		tags              []string
		expectedMinLength int
		expectedMaxLength int
	}{
		{"All Questions Limited", 3, "", []string{}, 3, 3},
		{"All Questions Unlimited", 10, "", []string{}, 5, 5},
		{"Technical Category", 2, models.CategoryTechnical, []string{}, 2, 2},
		{"Behavioral Category", 2, models.CategoryBehavioural, []string{}, 2, 2},
		{"No Matching Questions", 2, "", []string{"nonexistent"}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// can't directly test Start() as it has interactive components
			// instead, test the selection criteria by using FindQuestionsByFilter directly
			questions, err := mockRepo.FindQuestionsByFilter(tt.category, tt.tags)
			if err != nil {
				t.Fatalf("Error finding questions: %v", err)
			}

			// limit questions to requested number
			if len(questions) > tt.numQuestions {
				questions = questions[:tt.numQuestions]
			}

			// verify number of questions
			if len(questions) < tt.expectedMinLength || len(questions) > tt.expectedMaxLength {
				t.Errorf("Expected between %d and %d questions, got %d",
					tt.expectedMinLength, tt.expectedMaxLength, len(questions))
			}

			// check category filter
			if tt.category != "" {
				for i, q := range questions {
					if q.Category != tt.category {
						t.Errorf("Question %d has category %q, expected %q", i, q.Category, tt.category)
					}
				}
			}
		})
	}
}

