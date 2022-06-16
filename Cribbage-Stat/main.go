package main

import "fmt"

type card struct {
	Type  string
	Suit  string
	Value int
}

type cardSet []card

func createDeck() (deck cardSet) {

	deck = append(deck, card{Type: "Eight", Suit: "Heart", Value: 8})
	deck = append(deck, card{Type: "Seven", Suit: "Heart", Value: 7})
	return
}

type scoreCombo interface {
	Score(cards *cardSet)
}

func (self *cardSet) Get15Score() int {
	total := 0
	for _, card := range *self {
		total += card.Value
	}
	return total
}

func main() {
	deck := createDeck()

	fmt.Println(deck.Get15Score())

}
