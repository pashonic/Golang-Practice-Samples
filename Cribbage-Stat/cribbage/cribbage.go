package cribbage

import (
	"math/rand"
	"sort"
)

// Rank represents a card rank
type Rank int

// Rank represents a card rank
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

// Suit represents a card suit
type Suit int

// Suit represents a card suit
const (
	HEART Suit = iota
	SPADE
	DIAMOND
	CLUB
)

// Card represets a playing card
type Card struct {
	Rank  Rank
	Suit  Suit
	Value int
}

// CardSet represets a set of playing cards
type CardSet []Card

// CreateDeck returns a full standard deck of cards
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

// GetBestCards returns highest scoring set of given count of cards
func (cards *CardSet) GetBestCards(count int) CardSet {
	bestScore := 0
	var bestSet CardSet
	for i := 0; i <= len(*cards)-count; i++ {
		set, score := getCardCombinations(cards, count, CardSet{(*cards)[i]}, i+1)
		if bestScore <= score {
			bestSet = set
			bestScore = score
		}
	}
	return bestSet
}

func getCardCombinations(cards *CardSet, count int, set CardSet, start int) (CardSet, int) {
	if len(set) == count {
		setCopy := make(CardSet, len(set))
		copy(setCopy, set)
		return setCopy, setCopy.GetTotalScore()
	}

	bestScore := 0
	var bestSet CardSet
	for i := start; i < len(*cards); i++ {
		curSet, curScore := getCardCombinations(cards, count, append(set, (*cards)[i]), i+1)
		if bestScore <= curScore {
			bestSet = curSet
			bestScore = curScore
		}
	}
	return bestSet, bestScore
}

// Deal returns a given amount (count) of cards
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

// GetFifteenScore returns 15 score
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

// GetPairScore returns cribbage pair score
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

// GetRunScore returns cribbage run score
func (cards *CardSet) GetRunScore() int {

	// Copy cards and order copy.
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

	runTotal := 0
	runTracker := 1
	groupCounter := 1
	multiplier := 1
	var prevCard *Card = &copyCardSet[0]
	for i := 1; i < len(copyCardSet); i++ {

		// Handle different but increment rank card.
		if copyCardSet[i].Rank-1 == prevCard.Rank {
			runTracker++
			multiplier = multiplier * groupCounter
			groupCounter = 1

			// Handle same card.
		} else if copyCardSet[i].Rank == prevCard.Rank {
			groupCounter++

			// Finished this run, add to total, reset trackers.
		} else {
			runTotal += getRunTotal(runTracker, multiplier)
			groupCounter = 1
			multiplier = 1
			runTracker = 1
		}

		prevCard = &copyCardSet[i]
	}

	return runTotal + getRunTotal(runTracker, (multiplier*groupCounter))
}

// GetFlushScore returns flush score of given cards
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

// GetNobScore returns nob score with given cut card
func (cards *CardSet) GetNobScore(cutCard *Card) int {
	for _, card := range *cards {
		if card == *cutCard {
			continue
		}
		if card.Suit == cutCard.Suit {
			return 1
		}
	}
	return 0
}

// GetTotalScoreWithNob returns total score with given cut card
func (cards *CardSet) GetTotalScoreWithNob(cutCard *Card) int {
	return cards.GetTotalScore() + cards.GetNobScore(cutCard)
}

// GetTotalScore returns total score
func (cards *CardSet) GetTotalScore() int {
	scoreFunctions := []func() int{
		cards.GetFifteenScore,
		cards.GetFlushScore,
		cards.GetPairScore,
		cards.GetRunScore,
	}

	total := 0
	for _, scoreFunc := range scoreFunctions {
		total += scoreFunc()
	}
	return total
}
