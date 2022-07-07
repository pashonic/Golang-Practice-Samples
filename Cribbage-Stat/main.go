package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cribbage"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := cribbage.CreateDeck()

	fmt.Println(deck)

	fmt.Println(deck.GetPairScore())
}
