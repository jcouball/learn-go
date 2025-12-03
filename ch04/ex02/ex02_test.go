package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGenerateRandomInts(t *testing.T) {
	result := generateRandomInts()

	// Test that we have exactly 100 numbers
	if len(result) != 100 {
		t.Errorf("Expected 100 numbers, got %d", len(result))
	}

	// Test that each number is in the range [0, 99]
	for i, num := range result {
		if num < 0 || num > 99 {
			t.Errorf("Number at index %d is %d, expected value between 0 and 99 inclusive", i, num)
		}
	}
}

func TestGenerateDivisibilityStrings(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []string
	}{
		{
			name:     "divisible by 2 and 3",
			input:    []int{6, 12, 18},
			expected: []string{"Six!", "Six!", "Six!"},
		},
		{
			name:     "divisible by 2 only",
			input:    []int{2, 4, 8},
			expected: []string{"Two!", "Two!", "Two!"},
		},
		{
			name:     "divisible by 3 only",
			input:    []int{3, 9, 15},
			expected: []string{"Three!", "Three!", "Three!"},
		},
		{
			name:     "not divisible by 2 or 3",
			input:    []int{1, 5, 7},
			expected: []string{"Never mind", "Never mind", "Never mind"},
		},
		{
			name:     "mixed values",
			input:    []int{6, 2, 3, 5},
			expected: []string{"Six!", "Two!", "Three!", "Never mind"},
		},
		{
			name:     "zero is divisible by both",
			input:    []int{0},
			expected: []string{"Six!"},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateDivisibilityStrings(tc.input)

			if len(result) != len(tc.expected) {
				t.Errorf("Expected %d strings, got %d", len(tc.expected), len(result))
				return
			}

			for i, str := range result {
				if str != tc.expected[i] {
					t.Errorf("At index %d: expected %q, got %q", i, tc.expected[i], str)
				}
			}
		})
	}
}

func TestMain(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main
	main()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output format: should be a slice of strings
	if !strings.HasPrefix(output, "[") || !strings.HasSuffix(strings.TrimSpace(output), "]") {
		t.Errorf("Expected output to be a slice format [...]")
	}

	// Verify output contains expected strings (can't predict exact output due to randomness)
	if !strings.Contains(output, "Six!") && !strings.Contains(output, "Two!") && !strings.Contains(output, "Three!") && !strings.Contains(output, "Never mind") {
		t.Error("Expected output to contain at least one of: Six!, Two!, Three!, or Never mind")
	}
}
