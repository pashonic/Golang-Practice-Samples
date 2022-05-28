package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(seed)
	maxNumber := 100
	randomNumber := randomGen.Intn(maxNumber)

	fmt.Printf("Guess a number between 0 and %d\n", maxNumber)

	guessesLeft := 10
	var guessedNumber int
	for 0 < guessesLeft {
		fmt.Println("Guesses Left:", guessesLeft)

		// Get user input
		fmt.Scanln(&guessedNumber)

		if guessedNumber == randomNumber {
			break
		}

		if guessedNumber > randomNumber {
			fmt.Println("TOO HIGH")
		} else {
			fmt.Println("TOO LOW")
		}

		guessesLeft--
	}

	if guessesLeft == 0 {
		fmt.Println("!!!GAME OVER!!!")
	} else {
		fmt.Println("!!!WINNER!!!, number was ", randomNumber)
	}
}
