package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pashonic/Golang-Practice-Samples/Number-Guess/Version2/guessgame"
)

func processUserInput(gg *guessgame.GuessGame, guess uint64) bool {
	switch gg.Guess(guess) {
	case guessgame.WIN:
		fmt.Printf("!!!WINNER!!! with %v guesses left\n", gg.GetGuesses()-1)
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
	var randomNumber uint64 = randomGen.Uint64()
	var guesses uint64 = 10
	var maxValue uint64 = 100
	var guessGame = guessgame.New(guesses, randomNumber)

	for {
		fmt.Printf("Guesses Left: %v, enter guess between 0 and %v:", guessGame.GetGuesses(), maxValue)

		// Get user input
		var guessedNumber uint64
		fmt.Scanln(&guessedNumber)

		if !processUserInput(guessGame, guessedNumber) {
			break
		}
	}
}
