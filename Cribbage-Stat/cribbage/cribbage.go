package cribbage

import "math/rand"

type Rank int

const (
	ACE Rank = iota
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

type Suit int

const (
	HEART Suit = iota
	SPADE
	DIAMOND
	CLUB
)

type Card struct {
	Rank  Rank
	Suit  Suit
	value int
}

type CardSet []Card

func CreateDeck() (deck CardSet) {

	values := map[Rank]int{
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

	suits := []Suit{
		HEART,
		SPADE,
		DIAMOND,
		CLUB,
	}

	for rank, value := range values {
		for _, suit := range suits {
			deck = append(deck, Card{Rank: rank, Suit: suit, value: value})
		}
	}

	return deck
}

func (cards *CardSet) Deal(count int) CardSet {
	var outCards CardSet
	for i, _ := range rand.Perm(len(*cards))[:count] {
		outCards = append(outCards, (*cards)[i])
		*cards = append((*cards)[:i], (*cards)[i+1:]...)
	}
	return outCards
}

func (cards *CardSet) Get15Score() int {
	valueSet := []int{}
	for _, v := range *cards {
		valueSet = append(valueSet, v.value)
	}

	total := 0
	target := 15
	for i := 0; i < len(valueSet)-1; i++ {
		total += subsetSum(valueSet, target, []int{valueSet[i]}, i+1, valueSet[i])
	}

	return total * 2
}

func subsetSum(array []int, target int, pass []int, start int, sum int) int {
	if sum == target {
		return 1
	}

	if sum > target {
		return 0
	}

	total := 0
	for i := start; i < len(array); i++ {
		total += subsetSum(array, target, append(pass, array[i]), i+1, sum+array[i])
	}
	return total
}
