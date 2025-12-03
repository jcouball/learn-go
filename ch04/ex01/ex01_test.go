package main

import (
	"bytes"
	"fmt"
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

	// Verify output format: should be a slice of integers
	if !strings.HasPrefix(output, "[") || !strings.HasSuffix(strings.TrimSpace(output), "]") {
		t.Errorf("Expected output to be a slice format [...]")
	}

	// Parse and validate the output contains 100 numbers
	output = strings.TrimSpace(output)
	output = strings.TrimPrefix(output, "[")
	output = strings.TrimSuffix(output, "]")
	numbers := strings.Fields(output)

	if len(numbers) != 100 {
		t.Errorf("Expected 100 numbers in output, got %d", len(numbers))
	}

	// Verify each value is a valid number (just check we can print them)
	for _, numStr := range numbers {
		var num int
		if _, err := fmt.Sscanf(numStr, "%d", &num); err != nil {
			t.Errorf("Invalid number in output: %s", numStr)
		}
	}
}
