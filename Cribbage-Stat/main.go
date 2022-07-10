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
	var hand24 cribbage.CardSet = cribbage.CardSet{
		cribbage.Card{Rank: cribbage.EIGHT, Suit: cribbage.HEART, Value: 8},
		cribbage.Card{Rank: cribbage.SEVEN, Suit: cribbage.DIAMOND, Value: 7},
		cribbage.Card{Rank: cribbage.SEVEN, Suit: cribbage.SPADE, Value: 7},
		cribbage.Card{Rank: cribbage.SIX, Suit: cribbage.CLUB, Value: 6},
		cribbage.Card{Rank: cribbage.EIGHT, Suit: cribbage.CLUB, Value: 8},
	}
	fmt.Println(hand24.GetRunScore())

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
