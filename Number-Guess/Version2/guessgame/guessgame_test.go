package guessgame

import (
	"testing"
)

func TestResponses(t *testing.T) {
	var randomNumber uint64 = 500

	var guessGame = GuessGame{
		Guesses:      1001,
		RandomNumber: 500,
	}

	for i := uint64(0); i <= 1000; i++ {

		result := guessGame.Guess(i)

		if i == randomNumber {
			if result != WIN {
				t.Fatalf(`%v should of been a win`, i)
			}
		} else if i < randomNumber {
			if result != TOOLOW {
				t.Fatalf(`%v should of been low`, i)
			}
		} else {
			if result != TOOHIGH {
				t.Fatalf(`%v should of been high`, i)
			}
		}
	}

	if guessGame.Guess(randomNumber-1) != LOSS {
		t.Fatalf(`should of been LOSS`)
	}

	if guessGame.Guess(randomNumber+1) != GAMEOVER {
		t.Fatalf(`should of been GAMEOVER`)
	}
}
