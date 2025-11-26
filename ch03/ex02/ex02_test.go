package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMessageVariable(t *testing.T) {
	message := "Hi 👩 and and 👨"
	expected := "Hi 👩 and and 👨"

	if message != expected {
		t.Errorf("Expected message = %q, got %q", expected, message)
	}
}

func TestFourthRune(t *testing.T) {
	message := "Hi 👩 and and 👨"
	runes := []rune(message)

	// The fourth rune (index 3) should be '👩'
	if len(runes) < 4 {
		t.Fatalf("Expected at least 4 runes, got %d", len(runes))
	}

	fourthRune := runes[3]
	expected := '👩'

	if fourthRune != expected {
		t.Errorf("Expected fourth rune to be %c (U+%04X), got %c (U+%04X)",
			expected, expected, fourthRune, fourthRune)
	}
}

func TestRuneConversion(t *testing.T) {
	// Test that the message is correctly converted to runes
	message := "Hi 👩 and and 👨"
	runes := []rune(message)

	// Expected runes: 'H', 'i', ' ', '👩', ' ', 'a', 'n', 'd', ' ', 'a', 'n', 'd', ' ', '👨'
	expectedLength := 14

	if len(runes) != expectedLength {
		t.Errorf("Expected %d runes, got %d", expectedLength, len(runes))
	}
}

func TestMainOutput(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check that output contains the message and fourth rune
	expectedStrings := []string{
		"message:",
		"Hi 👩 and and 👨",
		"Fourth rune:",
		"👩",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain %q, but it didn't.\nOutput: %s", expected, output)
		}
	}
}
