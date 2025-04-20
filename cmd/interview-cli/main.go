package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"interview-cli/internal/config"
	"interview-cli/internal/repository"
	"interview-cli/internal/session"
)

func main() {

	// define command line flags
	addCmd := flag.Bool("add", false, "Add a new interview question")
	listCmd := flag.Bool("list", false, "List all interview questions")
	practiceCmd := flag.Bool("practice", false, "Start a practice session")
	numQuestions := flag.Int("n", 5, "Number of questions for practice session")
	categoryFlag := flag.String("category", "", "Filter questions by category (behavioral/technical)")
	tagsFlag := flag.String("tags", "", "Filter questions by tags (comma-separated)")

	flag.Parse()

	// parse tags if provided
	var tags []string
	if *tagsFlag != "" {
		for _, tag := range strings.Split(*tagsFlag, ",") {
			tags = append(tags, strings.TrimSpace(tag))
		}
	}

	// initialize configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// initialize repository
	repo := repository.NewQuestionRepository(cfg.QuestionsFile)

	var cmdErr error
	switch {
	case *addCmd:
		cmdErr = repo.AddQuestion()
	case *listCmd:
		cmdErr = repo.ListQuestions()
	case *practiceCmd:
		practice := session.NewPracticeSession(repo)
		cmdErr = practice.Start(*numQuestions, *categoryFlag, tags)
	default:
		// if no command specified, default to practice
		practice := session.NewPracticeSession(repo)
		cmdErr = practice.Start(*numQuestions, *categoryFlag, tags)
	}

	if cmdErr != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", cmdErr)
		os.Exit(1)
	}
}
