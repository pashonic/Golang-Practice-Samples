package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GuessGame struct {
	guesses      int
	maxValue     int
	randomNumber int
}

const (
	TOOHIGH = iota
	TOOLOW
	WON
	LOSS
)

func (gg *GuessGame) Guess(guess int) int {
	if guess == gg.randomNumber {
		return WON
	}

	gg.guesses--

	if gg.guesses < 1 {
		return LOSS
	}

	if guess > gg.randomNumber {
		return TOOHIGH
	} else {
		return TOOLOW
	}
}

func main() {

	// Seed the random generator or we get same value everytime
	seed := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(seed)

	maxValue := 100
	guesses := 10
	randomNumber := randomGen.Intn(maxValue)

	var guessGame = GuessGame{
		guesses:      guesses,
		maxValue:     maxValue,
		randomNumber: randomNumber,
	}

gameLoop:
	for {
		fmt.Printf("Guesses Left: %v, enter guess between 0 and %v:", guessGame.guesses, maxValue)

		// Get user input
		var guessedNumber int
		fmt.Scanln(&guessedNumber)

		switch guessGame.Guess(guessedNumber) {
		case WON:
			fmt.Printf("!!!WINNER!!! with %v guesses left\n", guessGame.guesses)
			break gameLoop
		case LOSS:
			fmt.Println("!!!GAME OVER!!!")
			break gameLoop
		case TOOHIGH:
			fmt.Println("TOO HIGH")
		case TOOLOW:
			fmt.Println("TOO LOW")
		}
	}
}
