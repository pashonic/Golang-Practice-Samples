package guessgame

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	var guessGame = GuessGame{
		Guesses:      10,
		MaxValue:     100,
		RandomNumber: 45,
	}

	// BUGBUG, add more.
	if guessGame.Guess(20) != TOOLOW {
		t.Fatalf("Should be TOO LOW")
	}
}
