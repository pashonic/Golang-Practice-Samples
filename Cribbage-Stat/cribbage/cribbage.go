package cribbage

import (
	"fmt"
	"math/rand"
	"sort"
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
	Value int
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
			deck = append(deck, Card{Rank: rank, Suit: suit, Value: value})
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
		valueSet = append(valueSet, card.Value)
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
		6: -1,
	}

	total := 0
	for _, card := range *cards {
		if subTotals[int(card.Rank)] == -1 {
			continue
		}
		total += subTotals[int(card.Rank)]
		subTotals[int(card.Rank)] = values[subTotals[int(card.Rank)]]
	}

	return total
}

func cardComp(cardA, cardB Card) bool {
	return int(cardA.Rank) > int(cardB.Rank)
}

func (cards *CardSet) GetRunScore() int {
	copyCardSet := make(CardSet, len(*cards))
	copy(copyCardSet, *cards)
	sort.Slice(copyCardSet, func(i, j int) bool {
		return (copyCardSet)[i].Rank < (copyCardSet)[j].Rank
	})

	scoreTotal := 0
	runTracker := 0
	var prevCard *Card = nil
	fmt.Println(copyCardSet)
	for i := 0; i < len(copyCardSet); i++ {
		if prevCard != nil {
			if copyCardSet[i].Rank-1 == prevCard.Rank {
				runTracker++
				if runTracker == 3 {
					scoreTotal += runTracker
				} else if runTracker > 3 {
					scoreTotal++
				}
			} else {
				runTracker = 1
			}
		} else {
			runTracker++
		}
		prevCard = &copyCardSet[i]
	}

	return scoreTotal
}
