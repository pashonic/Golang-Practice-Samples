package guessgame

import (
	"testing"
)

func TestResponses(t *testing.T) {
	var maxValue uint64 = 1000
	var randomNumber uint64 = maxValue / 2
	var guessGame = New(maxValue+1, randomNumber)

	for i := uint64(0); i <= maxValue; i++ {

		result := guessGame.Guess(i)

		if i == randomNumber {
			if result != WIN {
				t.Fatalf(`%v should of been a WIN`, i)
			}
		} else if i < randomNumber {
			if result != TOOLOW {
				t.Fatalf(`%v should of been LOW`, i)
			}
		} else {
			if result != TOOHIGH {
				t.Fatalf(`%v should of been HIGH`, i)
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
