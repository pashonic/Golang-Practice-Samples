package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cardset"
	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cribbage"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := cribbage.CreateDeck()

	hand := cardset.Deal(deck, 4)

	fmt.Println(deck)
	fmt.Println("BUGBUG")
	fmt.Println(hand)
}
