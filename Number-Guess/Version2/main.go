package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pashonic/Golang-Practice-Samples/Number-Guess/Version2/guessgame"
)

func processUserInput(gg *guessgame.GuessGame, guess int) bool {
	switch gg.Guess(guess) {
	case guessgame.WON:
		fmt.Printf("!!!WINNER!!! with %v guesses left\n", gg.Guesses-1)
		return false
	case guessgame.LOSS:
		fmt.Println("!!!GAME OVER!!!")
		return false
	case guessgame.TOOHIGH:
		fmt.Println("TOO HIGH")
	case guessgame.TOOLOW:
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

	var guessGame = guessgame.GuessGame{
		Guesses:      guesses,
		MaxValue:     maxValue,
		RandomNumber: randomNumber,
	}

	for {
		fmt.Printf("Guesses Left: %v, enter guess between 0 and %v:", guessGame.Guesses, maxValue)

		// Get user input
		var guessedNumber int
		fmt.Scanln(&guessedNumber)

		if !processUserInput(&guessGame, guessedNumber) {
			break
		}
	}
}
