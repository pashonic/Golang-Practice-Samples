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
		cribbage.Card{Rank: cribbage.ACE, Suit: cribbage.SPADE, Value: 6},
		cribbage.Card{Rank: cribbage.TWO, Suit: cribbage.HEART, Value: 6},
		cribbage.Card{Rank: cribbage.THREE, Suit: cribbage.CLUB, Value: 5},
		cribbage.Card{Rank: cribbage.THREE, Suit: cribbage.CLUB, Value: 5},
		cribbage.Card{Rank: cribbage.TEN, Suit: cribbage.SPADE, Value: 6},
		cribbage.Card{Rank: cribbage.JOKER, Suit: cribbage.HEART, Value: 6},
		cribbage.Card{Rank: cribbage.QUEEN, Suit: cribbage.CLUB, Value: 5},
		cribbage.Card{Rank: cribbage.QUEEN, Suit: cribbage.CLUB, Value: 5},
	}

	fmt.Println(deck.GetRunScore())

}
