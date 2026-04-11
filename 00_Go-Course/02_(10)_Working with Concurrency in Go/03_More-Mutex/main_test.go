package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	// Save original stdout
	stdOut := os.Stdout
	defer func() { os.Stdout = stdOut }()

	// Create pipe
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe error: %v", err)
	}
	defer r.Close()

	// Redirect stdout
	os.Stdout = w

	// Run main
	main()

	// Close writer so reader gets EOF
	if err := w.Close(); err != nil {
		t.Fatalf("close error: %v", err)
	}

	// Read output
	result, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("read error: %v", err)
	}

	output := string(result)

	// Assertion
	if !strings.Contains(output, "$3430.00") {
		t.Errorf("Wrong Balance returned, got: %s", output)
	}
}
