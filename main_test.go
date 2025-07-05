package main

import (
	"io"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func TestPrintRandomCatArt(t *testing.T) {
	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a deterministic random generator for testing
	rnd := rand.New(rand.NewSource(12345))

	// Call the function to test
	err := printRandomCatArt(rnd)

	// Restore stdout and read the captured output
	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)

	// Check for errors
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

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

	// Check if output contains ANSI escape codes for clearing screen
	if !strings.Contains(string(out), "\033[H\033[2J") {
		t.Errorf("Expected output to contain ANSI clear screen codes")
	}
}

func TestPrintRandomCatArtWithDifferentSeeds(t *testing.T) {
	// Test with different seeds to ensure randomness works
	seeds := []int64{1, 2, 3, 4, 5}
	results := make([]string, len(seeds))

	for i, seed := range seeds {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		rnd := rand.New(rand.NewSource(seed))
		err := printRandomCatArt(rnd)

		w.Close()
		os.Stdout = old
		out, _ := io.ReadAll(r)

		if err != nil {
			t.Errorf("Expected no error for seed %d, but got: %v", seed, err)
		}

		results[i] = string(out)
	}

	// Check that we get some variation in outputs (not all identical)
	allSame := true
	for i := 1; i < len(results); i++ {
		if results[i] != results[0] {
			allSame = false
			break
		}
	}

	if allSame {
		t.Errorf("Expected some variation in outputs with different seeds")
	}
}

func TestConstants(t *testing.T) {
	// Test that constants are defined correctly
	if AnimationLoops <= 0 {
		t.Errorf("AnimationLoops should be positive, got %d", AnimationLoops)
	}

	if FrameDelay <= 0 {
		t.Errorf("FrameDelay should be positive, got %v", FrameDelay)
	}

	// Test that FrameDelay is reasonable (between 1ms and 1s)
	if FrameDelay < time.Millisecond || FrameDelay > time.Second {
		t.Errorf("FrameDelay should be between 1ms and 1s, got %v", FrameDelay)
	}
}

func TestCatArtsValid(t *testing.T) {
	// Test that catArts is properly defined
	if len(catArts) == 0 {
		t.Errorf("catArts should not be empty")
	}

	for i, catPair := range catArts {
		if len(catPair) == 0 {
			t.Errorf("catPair at index %d should not be empty", i)
		}
		for j, catArt := range catPair {
			if strings.TrimSpace(catArt) == "" {
				t.Errorf("catArt at index [%d][%d] should not be empty", i, j)
			}
		}
	}
}
