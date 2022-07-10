package main

import (
	"fmt"

	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cribbage"
)

func main() {

	// Create a deck of cards
	deck := cribbage.CreateDeck()

	// Deal 6 cards
	dealtHand := deck.Deal(6)

	// Get the best 4 cards that result in highest score
	selectedHand := dealtHand.GetBestCards(4)

	// Deal cut card
	var cutCard cribbage.Card = deck.Deal(1)[0]

	// Add cut card to deck
	dealtHand = append(selectedHand, cutCard)

	// Get Scores
	fmt.Println(dealtHand)
	fmt.Println(cutCard)
	fmt.Println(dealtHand.GetTotalScoreWithNob(&cutCard))

}
