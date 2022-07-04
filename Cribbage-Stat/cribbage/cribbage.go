package cribbage

import (
	"github.com/pashonic/Golang-Practice-Samples/Cribbage-Stat/cardset"
)

type CribCard struct {
	cardset.Card
	valued int
}

func CreateDeck() (deck []cardset.CardInterface) {

	values := map[cardset.Rank]int{
		cardset.ACE:   1,
		cardset.TWO:   2,
		cardset.THREE: 3,
		cardset.FOUR:  4,
		cardset.FIVE:  5,
		cardset.SIX:   6,
		cardset.SEVEN: 7,
		cardset.EIGHT: 8,
		cardset.NINE:  9,
		cardset.TEN:   10,
		cardset.JOKER: 10,
		cardset.QUEEN: 10,
		cardset.KING:  10,
	}

	suits := []cardset.Suit{
		cardset.HEART,
		cardset.SPADE,
		cardset.DIAMOND,
		cardset.CLUB,
	}

	for rank, value := range values {
		for _, suit := range suits {
			deck = append(deck, CribCard{cardset.Card{Rank: rank, Suit: suit}, value})
		}
	}

	return deck
}

/*
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
*/
