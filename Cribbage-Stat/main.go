package main

import "github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cribbage"

func main() {

	deck := cribbage.CreateDeck()

	deck.GetRunScore()
}
