package cribbage

import (
	"math/rand"
)

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
	len := len(*cards)
	for j := 0; j < count; j++ {
		index := rand.Intn(len)
		outCards = append(outCards, (*cards)[index])
		*cards = append((*cards)[:index], (*cards)[index+1:]...)
		len--
	}
	return outCards
}

func (cards *CardSet) GetFifteenScore() int {
	valueSet := []int{}
	for _, card := range *cards {
		valueSet = append(valueSet, card.value)
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

func (cards *CardSet) GetPairScore() int {
	subTotals := map[int]int{}
	values := map[int]int{
		0: 2,
		2: 4,
		4: 6,
		6: 12,
	}

	total := 0
	for _, card := range *cards {
		if subTotals[int(card.Rank)] > 6 {
			continue
		}
		total += subTotals[int(card.Rank)]
		subTotals[int(card.Rank)] = values[subTotals[int(card.Rank)]]
	}

	return total
}
