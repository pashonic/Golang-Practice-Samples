package main

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	var guessGame = GuessGame{
		guesses:      10,
		maxValue:     100,
		randomNumber: 45,
	}

	// BUGBUG, add more.
	if guessGame.guess(20) != TOOLOW {
		t.Fatalf("Should be TOO LOW")
	}

}
