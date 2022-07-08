package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cribbage"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var deck = cribbage.CardSet{
		{Rank: cribbage.ACE, Suit: cribbage.CLUB, Value: 3},
		{Rank: cribbage.TWO, Suit: cribbage.CLUB, Value: 3},
		{Rank: cribbage.THREE, Suit: cribbage.CLUB, Value: 3},
		{Rank: cribbage.FOUR, Suit: cribbage.CLUB, Value: 3},
	}

	fmt.Println(deck.GetRunScore())
}
