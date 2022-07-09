package main

import (
	"fmt"

	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cribbage"
)

func main() {
	/*
		deck := cribbage.CreateDeck()

		hand := deck.Deal(6)

		hand.GetBestCards(4)

	*/
	deck := cribbage.CardSet{
		cribbage.Card{Rank: cribbage.FOUR, Suit: cribbage.HEART, Value: 7},
		cribbage.Card{Rank: cribbage.FOUR, Suit: cribbage.HEART, Value: 8},
		cribbage.Card{Rank: cribbage.FOUR, Suit: cribbage.HEART, Value: 7},
	}

	fmt.Println(deck.GetPairScore())

	/*
		deck := cribbage.CardSet{
			cribbage.Card{Rank: cribbage.SEVEN, Suit: cribbage.SPADE, Value: 7},
			cribbage.Card{Rank: cribbage.EIGHT, Suit: cribbage.HEART, Value: 8},
			cribbage.Card{Rank: cribbage.SEVEN, Suit: cribbage.HEART, Value: 7},
			cribbage.Card{Rank: cribbage.FOUR, Suit: cribbage.HEART, Value: 4},
		}

		fmt.Println(deck.GetBestCards(4))
	*/
}
