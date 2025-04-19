package models

import (
	"slices"
)

const (
	CategoryBehavioural = "behavioural"
	CategoryTechnical   = "technical"
)

var ValidCategories = []string{CategoryBehavioural, CategoryTechnical}

// Question → interview question w/ metadata
type Question struct {
	Text     string   `json:"text"`
	Category string   `json:"category"`
	Tags     []string `json:"tags,omitempty"`
}

// checks if category is valid
func isValidCategory(category string) bool {
	return slices.Contains(ValidCategories, category)
}

// checks if question has specific tag
func (q *Question) HasTag(tag string) bool {
	return slices.Contains(q.Tags, tag)
}

// checks if question has all the provided tags
func (q *Question) HasAllTags(tags []string) bool {
	if len(tags) == 0 {
		return true
	}

	for _, tag := range tags {
		if !q.HasTag(tag) {
			return false
		}
	}
	return true
}
