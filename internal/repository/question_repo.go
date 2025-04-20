package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"interview-cli/internal/models"
)

// handles storage and retrieval of interview questions
type QuestionRepository struct {
	filePath string // path to json file storing questions
}

// creates new repo with specified path
func NewQuestionRepository(filePath string) *QuestionRepository {
	return &QuestionRepository{
		filePath: filePath,
	}
}

// reads all questions from JSON file and returns qeustions slice + errors
func (r *QuestionRepository) LoadQuestions() ([]models.Question, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read questions file: %w", err)
	}

	var questions []models.Question
	if err := json.Unmarshal(data, &questions); err != nil {
		return nil, fmt.Errorf("failed to parse questions JSON: %w", err)
	}

	return questions, nil
}

// writes the questions slice to the JSON file
func (r *QuestionRepository) SaveQuestions(questions []models.Question) error {
	data, err := json.MarshalIndent(questions, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal questions to JSON: %w", err)
	}
	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write questions to file: %w", err)
	}

	return nil
}

// prompt usere to input new question w/ category + tags
// add question to repo and saves updated list
func (r *QuestionRepository) AddQuestions() error {
	// load existing questions, continue if file doesn't exist yet
	questions, err := r.LoadQuestions()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	questionText := getInput("Enter the interview questions: ")
	if questionText == "" {
		return fmt.Errorf("question text cannot be empty")
	}

	// get category w/ validation
	var category string

	for {
		category = getInput(fmt.Sprintf("Enter category (%s/%s): ", models.CategoryBehavioural, models.CategoryTechnical))

		category = strings.ToLower(strings.TrimSpace(category))

		if models.IsValidCategory(category) {
			break
		}

		fmt.Printf("Invalid category. Please enter either '%s' or '%s'.\n", models.CategoryBehavioural, models.CategoryTechnical)
	}

	// get tags as comma-separated values
	tagsInput := getInput("Enter tags (comma-separated): ")
	var tags []string
	if tagsInput != "" {
		for _, tag := range strings.Split(tagsInput, ",") {
			tags = append(tags, strings.TrimSpace(tag))
		}
	}

	// create and add new question
	newQuestion := models.Question{
		Text:     questionText,
		Category: category,
		Tags:     tags,
	}
	questions = append(questions, newQuestion)

	// save updated questions list
	if err := r.SaveQuestions(questions); err != nil {
		return err
	}

	fmt.Println("Question added successfully!")
	return nil
}

// displays all questions grouped by category
// show question count + associated tags
func (r *QuestionRepository) ListQuestions() error {
	questions, err := r.LoadQuestions()
	if err != nil {
		return err
	}

	// handle case where no questions exist
	if len(questions) == 0 {
		fmt.Println("No questions found. Add some questions first.")
		return nil
	}

	fmt.Printf("Total questions: %d\n\n", len(questions))

	// group questions by category
	categories := make(map[string][]models.Question)
	for _, q := range questions {
		categories[q.Category] = append(categories[q.Category], q)
	}

	// display questions grouped by category
	for category, qs := range categories {
		if category == "" {
			category = "Uncategorized"
		}
		fmt.Printf("Category: %s (%d questions)\n", category, len(qs))
		for i, q := range qs {
			// format tags if they exit
			tags := ""
			if len(q.Tags) > 0 {
				tags = fmt.Sprintf(" [%s]", strings.Join(q.Tags, ", "))
			}
			fmt.Printf("  %d. %s%s\n", i+1, q.Text, tags)
		}
		fmt.Println()
	}

	return nil
}

// return questions matching specified category and tags
// if no filters → all questions
func (r *QuestionRepository) FindQuestionsByFilter(category string, tags []string) ([]models.Question, error) {
	questions, err := r.LoadQuestions()
	if err != nil {
		return nil, err
	}

	// if no filter provided, return all qs
	if category == "" && len(tags) == 0 {
		return questions, nil
	}

	// apply filters to find matching qs
	var filtered []models.Question
	for _, q := range questions {
		// skip if category doesn't match
		if category != "" && q.Category != category {
			continue
		}

		// skip if question doesn't have specified tags
		if !q.HasAllTags(tags) {
			continue
		}

		filtered = append(filtered, q)
	}

	return filtered, nil
}

// prompt user and returns trimmed input string
func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
