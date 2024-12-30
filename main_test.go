package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintRandomCatArt(t *testing.T) {
	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function to test
	printRandomCatArt()

	// Restore stdout and read the captured output
	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)

	// Check if the output contains any of the cat arts
	found := false
	for _, catPair := range catArts {
		for _, catArt := range catPair {
			if strings.Contains(string(out), catArt) {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		t.Errorf("Expected output to contain one of the cat arts, but it didn't")
	}
}
