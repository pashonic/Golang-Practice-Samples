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

func ProcessUserInput(gg *GuessGame, guess int) bool {
	switch gg.Guess(guess) {
	case WON:
		fmt.Printf("!!!WINNER!!! with %v guesses left\n", gg.guesses-1)
		return false
	case LOSS:
		fmt.Println("!!!GAME OVER!!!")
		return false
	case TOOHIGH:
		fmt.Println("TOO HIGH")
	case TOOLOW:
		fmt.Println("TOO LOW")
	}
	return true
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

	for {
		fmt.Printf("Guesses Left: %v, enter guess between 0 and %v:", guessGame.guesses, maxValue)

		// Get user input
		var guessedNumber int
		fmt.Scanln(&guessedNumber)

		if !ProcessUserInput(&guessGame, guessedNumber) {
			break
		}

	}

}
