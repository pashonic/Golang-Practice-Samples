package main

import (
	"fmt"
	"math/rand"
	"time"
)

type rank int

const (
	ACE rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JOKER
	QUEEN
	KING
)

type suit int

const (
	HEART suit = iota
	SPADE
	DIAMOND
	CLUB
)

type card struct {
	rank  rank
	suit  suit
	value int
}
type cardSet []card

func (self cardSet) deal(count int) (cards cardSet) {
	for i, _ := range rand.Perm(len(self))[:count] {
		cards = append(cards, self[i])
		self = append(self[:i], self[i+1:]...)
	}
	return cards
}

func createDeck() (deck cardSet) {

	values := map[rank]int{
		ACE:   1,
		TWO:   2,
		THREE: 3,
		FOUR:  4,
		FIVE:  5,
		SIX:   6,
		SEVEN: 7,
		EIGHT: 8,
		NINE:  9,
		TEN:   10,
		JOKER: 10,
		QUEEN: 10,
		KING:  10,
	}

	suits := []suit{
		HEART, SPADE, DIAMOND, CLUB,
	}

	for rank, value := range values {
		for _, suit := range suits {
			deck = append(deck, card{rank: rank, suit: suit, value: value})
		}
	}

	return deck
}

func get15Score(cards cardSet) int {
	valueSet := []int{}
	for _, v := range cards {
		valueSet = append(valueSet, v.value)
	}
	return subsetSum(valueSet, 15) * 2
}

func subsetSum(array []int, target int) int {
	total := 0

	for i := 0; i < len(array)-1; i++ {
		total += subsetSumHelp(array, target, []int{array[i]}, i+1, array[i])
	}
	return total
}

func subsetSumHelp(array []int, target int, pass []int, start int, sum int) int {
	if sum == target {
		return 1
	}

	if sum > target {
		return 0
	}

	total := 0
	for i := start; i < len(array); i++ {
		total += subsetSumHelp(array, target, append(pass, array[i]), i+1, sum+array[i])
	}
	return total
}

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := createDeck()

	hand := deck.deal(5)

	fmt.Println(hand)
	fmt.Println(get15Score(hand))
}