package main

import "testing"

func TestProcessGuessCorrect(t *testing.T) {
	result := processGuess("thumb", "thumb")

	if result[0] != Correct {
		t.Fatalf("First letter must match")
	}

	if result[1] != Correct {
		t.Fatalf("Second letter must match")
	}

	if result[2] != Correct {
		t.Fatalf("Third letter must match")
	}

	if result[3] != Correct {
		t.Fatalf("Fourth letter must match")
	}
}

func TestTooLongInput(t *testing.T) {
	result := processGuess("wwwwwwwwww", "thumb")
	// should not panic
	if result[0] != Invalid {
		t.Fatalf("First letter must match")
	}
}

func TestTooShort(t *testing.T) {
	result := processGuess("w", "thumb")
	// should not panic
	if result[0] != Invalid {
		t.Fatalf("First letter must match")
	}
}

func TestProcessGuessContains(t *testing.T) {
	result := processGuess("oxllh", "hello")

	if result[0] != Contains {
		t.Fatalf("First should not be found but got %d", result[0])
	}

	if result[1] != Invalid {
		t.Fatalf("Second should not be found but got %d", result[0])

	}

	if result[2] != Correct {
		t.Fatalf("Third should not be found but got %d", result[0])

	}
}
