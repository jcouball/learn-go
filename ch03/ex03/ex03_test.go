package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

	expectedOutput := `emp1:  {John Doe 10001}
emp2:  {Jane Austin 10002}
emp3:  {Bob Marley 10003}
`

	if diff := cmp.Diff(expectedOutput, output); diff != "" {
		t.Errorf("Output mismatch (-want +got):\n%s", diff)
	}
}
