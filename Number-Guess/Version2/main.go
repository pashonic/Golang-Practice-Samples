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

type guessResult int

const (
	TOOHIGH guessResult = iota
	TOOLOW
	WON
	LOSS
)

func (self *GuessGame) guess(guess int) guessResult {
	if guess == self.randomNumber {
		return WON
	}

	self.guesses--

	if self.guesses < 1 {
		return LOSS
	}

	if guess > self.randomNumber {
		return TOOHIGH
	} else {
		return TOOLOW
	}
}

func processUserInput(gg *GuessGame, guess int) bool {
	switch gg.guess(guess) {
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

		if !processUserInput(&guessGame, guessedNumber) {
			break
		}

	}
}
