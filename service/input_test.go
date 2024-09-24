package service

import (
	"testing"
)

// Correctly assigns Words.name, Words.task, and Words.time based on input string
func TestCorrectAssignment(t *testing.T) {
	input := []string{"John", "Meeting", "10"}
	Input(input)

	if Words.name != "John" {
		t.Errorf("Expected name to be 'John', got %s", Words.name)
	}

	if Words.task != "Meeting" {
		t.Errorf("Expected task to be 'Meeting', got %s", Words.task)
	}
}
