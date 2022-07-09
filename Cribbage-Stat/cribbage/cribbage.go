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

func (cards *CardSet) GetRunScore() int {
	copyCardSet := make(CardSet, len(*cards))
	copy(copyCardSet, *cards)
	sort.Slice(copyCardSet, func(i, j int) bool {
		return (copyCardSet)[i].Rank < (copyCardSet)[j].Rank
	})

	getRunTotal := func(runTracker int, multiplier int) int {
		if runTracker >= 3 {
			return (runTracker * multiplier)
		}
		return 0
	}

	scoreTotal := 0
	runTracker := 1
	multiplier := 1
	var prevCard *Card = nil
	for i := 0; i < len(copyCardSet); i++ {
		if prevCard != nil {
			if copyCardSet[i].Rank-1 == prevCard.Rank {
				runTracker++
			} else if copyCardSet[i].Rank == prevCard.Rank {
				multiplier++
			} else {
				scoreTotal += getRunTotal(runTracker, multiplier)
				multiplier = 1
				runTracker = 1
			}
		}
		prevCard = &copyCardSet[i]
	}
	scoreTotal += getRunTotal(runTracker, multiplier)
	return scoreTotal
}

func (cards *CardSet) GetFlushScore() int {
	suitMap := map[Suit]int{
		HEART:   0,
		SPADE:   0,
		DIAMOND: 0,
		CLUB:    0,
	}

	for _, card := range *cards {
		suitMap[card.Suit]++
	}

	scoreTotal := 0
	for _, total := range suitMap {
		if total >= 4 {
			scoreTotal += total
		}
	}

	return scoreTotal
}

func (cards *CardSet) GetBestCards(count int) {
	for i := 0; i <= len(*cards)-count; i++ {
		set, score := getCombinations(cards, count, CardSet{(*cards)[i]}, i+1)
		fmt.Println(set, "-", score)
	}
}

func getCombinations(cards *CardSet, count int, set CardSet, start int) (CardSet, int) {
	if len(set) == count {
		setCopy := make(CardSet, len(set))
		copy(setCopy, set)
		return setCopy, setCopy.GetTotalScore()
	}

	bestScore := 0
	var bestSet CardSet
	for i := start; i < len(*cards); i++ {
		curSet, curScore := getCombinations(cards, count, append(set, (*cards)[i]), i+1)
		if bestScore <= curScore {
			bestSet = curSet
			bestScore = curScore
		}
	}
	return bestSet, bestScore
}

func (cards *CardSet) GetNobScore(cutCard *Card) int {
	for _, card := range *cards {
		if card.Rank == JOKER && card.Suit == cutCard.Suit {
			return 1
		}
	}
	return 0
}

func (cards *CardSet) GetTotalScore() int {
	scoreFunctions := []func() int{
		cards.GetFifteenScore,
		cards.GetFlushScore,
		cards.GetPairScore,
		cards.GetRunScore,
	}

	total := 0
	for _, scoreFun := range scoreFunctions {
		total += scoreFun()
	}
	return total
}
