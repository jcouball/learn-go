package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestGreetingsSlice(t *testing.T) {
	// Test that greetings slice is created with correct values
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	expected := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	if len(greetings) != len(expected) {
		t.Errorf("Expected greetings length %d, got %d", len(expected), len(greetings))
	}

	for i, greeting := range greetings {
		if greeting != expected[i] {
			t.Errorf("Expected greetings[%d] = %q, got %q", i, expected[i], greeting)
		}
	}
}

func TestSubslices(t *testing.T) {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	tests := []struct {
		name     string
		subslice []string
		expected []string
	}{
		{
			name:     "First two values",
			subslice: greetings[0:2],
			expected: []string{"Hello", "Hola"},
		},
		{
			name:     "Second, third, and fourth values",
			subslice: greetings[1:4],
			expected: []string{"Hola", "नमस्कार", "こんにちは"},
		},
		{
			name:     "Fourth and fifth values",
			subslice: greetings[3:5],
			expected: []string{"こんにちは", "Привіт"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.subslice) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(tt.subslice))
				return
			}

			for i := range tt.subslice {
				if tt.subslice[i] != tt.expected[i] {
					t.Errorf("Expected %q, got %q", tt.expected[i], tt.subslice[i])
				}
			}
		})
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

	// Check that output contains all greetings
	expectedStrings := []string{
		"greetings:",
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
		"subslice1:",
		"subslice2:",
		"subslice3:",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain %q, but it didn't.\nOutput: %s", expected, output)
		}
	}
}
