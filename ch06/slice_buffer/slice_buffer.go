package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func processData(data []byte) {
	// data shares the same underlying array as buffer
	// We can read from it but shouldn't modify it if we want to preserve the buffer
	fmt.Printf("Processing %d bytes: %q\n", len(data), string(data))
}

func main() {
	// Create a temporary file for demonstration
	content := "This is a sample file content that we will read using a reusable buffer."
	tmpFile, err := os.CreateTemp("", "buffer-example-*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	if _, err := tmpFile.WriteString(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	// Now read the file using a reusable buffer
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a reusable buffer once
	buffer := make([]byte, 16) // Small buffer to demonstrate multiple reads
	readCount := 0

	for {
		// Read into the existing buffer (no new allocation)
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading file: %v", err)
			}
			break
		}

		readCount++
		fmt.Printf("Read #%d: ", readCount)

		// Process the data in buffer[:bytesRead]
		processData(buffer[:bytesRead])

		// The buffer can be reused in the next iteration
	}

	// Alternative: using strings.Reader for demonstration
	fmt.Println("\n=== Reading from strings.Reader ===")
	reader := strings.NewReader("Hello, buffer reuse!")
	buffer2 := make([]byte, 8)
	readCount = 0

	for {
		bytesRead, err := reader.Read(buffer2)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading: %v", err)
			}
			break
		}

		readCount++
		fmt.Printf("Read #%d: ", readCount)
		processData(buffer2[:bytesRead])
	}
}
