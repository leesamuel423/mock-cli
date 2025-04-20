package session

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"interview-cli/internal/models"
	"interview-cli/internal/repository"
)

// handle the interview practice workflow
type PracticeSession struct {
	repo           *repository.QuestionRepository
	askedQuestions map[string]bool
}

// create new practice session
func NewPracticeSession(repo *repository.QuestionRepository) *PracticeSession {
	return &PracticeSession{
		repo:           repo,
		askedQuestions: make(map[string]bool),
	}
}

// begin a practice session w/ given parameters
func (p *PracticeSession) Start(numQuestions int, category string, tags []string) error {
	questions, err := p.repo.FindQuestionsByFilter(category, tags)
	if err != nil {
		return err
	}

	if len(questions) == 0 {
		fmt.Println("No questions found matching your criteria. Add some questions first or change your filters.")
		return nil
	}

	fmt.Println("\n=== INTERVIEW PRACTICE SESSION ===")
	fmt.Println("Press Enter after each question to continue, or type 'quit' to end the session.")
	fmt.Println()

	count := 0
	maxAttempts := len(questions) * 2 // To avoid infinite loop if all questions have been asked

	for count < numQuestions && maxAttempts > 0 {
		if len(questions) <= len(p.askedQuestions) {
			fmt.Println("You've gone through all available questions in this category.")
			break
		}

		q := p.getRandomQuestion(questions)
		if q == nil {
			maxAttempts--
			continue
		}

		count++
		fmt.Printf("\nQuestion %d/%d:\n", count, numQuestions)
		fmt.Printf("\033[1m%s\033[0m\n", q.Text)

		if category != "" || len(tags) > 0 {
			fmt.Printf("Category: %s", q.Category)
			if len(q.Tags) > 0 {
				fmt.Printf(" | Tags: %s", strings.Join(q.Tags, ", "))
			}
			fmt.Println()
		}

		input := getInput("\nPress Enter for next question or type 'quit' to end: ")
		if strings.ToLower(input) == "quit" {
			break
		}
	}

	fmt.Println("\nPractice session completed!")
	return nil
}

// return random question that hasn't been asked yet
func (p *PracticeSession) getRandomQuestion(questions []models.Question) *models.Question {
	if len(questions) == 0 {
		return nil
	}

	// create list of questions that haven't been asked yet
	var available []models.Question
	for _, q := range questions {
		if !p.askedQuestions[q.Text] {
			available = append(available, q)
		}
	}

	if len(available) == 0 {
		return nil
	}

	// select random question from available ones
	randIndex := rand.Intn(len(available))
	selectedQuestion := available[randIndex]

	// mark this question as asked
	p.askedQuestions[selectedQuestion.Text] = true

	return &selectedQuestion
}

// get user input from console
func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
