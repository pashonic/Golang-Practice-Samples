package cribbage

import (
	"testing"
)

var hand29CutCard = Card{Rank: FIVE, Suit: HEART, Value: 5}
var hand29 CardSet = CardSet{
	Card{Rank: JOKER, Suit: HEART, Value: 10},
	Card{Rank: FIVE, Suit: DIAMOND, Value: 5},
	Card{Rank: FIVE, Suit: SPADE, Value: 5},
	Card{Rank: FIVE, Suit: CLUB, Value: 5},
	hand29CutCard,
}

var handCutCard2 = Card{Rank: EIGHT, Suit: SPADE, Value: 8}
var nobHand CardSet = CardSet{
	Card{Rank: THREE, Suit: SPADE, Value: 3},
	Card{Rank: THREE, Suit: HEART, Value: 3},
	Card{Rank: TWO, Suit: SPADE, Value: 2},
	Card{Rank: TEN, Suit: DIAMOND, Value: 10},
	handCutCard2,
}

var handDoubleRun CardSet = CardSet{
	Card{Rank: EIGHT, Suit: HEART, Value: 8},
	Card{Rank: SEVEN, Suit: DIAMOND, Value: 7},
	Card{Rank: SEVEN, Suit: SPADE, Value: 7},
	Card{Rank: SIX, Suit: CLUB, Value: 6},
	Card{Rank: TEN, Suit: CLUB, Value: 10},
}

var hand24 CardSet = CardSet{
	Card{Rank: EIGHT, Suit: HEART, Value: 8},
	Card{Rank: SEVEN, Suit: DIAMOND, Value: 7},
	Card{Rank: SEVEN, Suit: SPADE, Value: 7},
	Card{Rank: SIX, Suit: CLUB, Value: 6},
	Card{Rank: EIGHT, Suit: CLUB, Value: 8},
}

var handTriplerun CardSet = CardSet{
	Card{Rank: JOKER, Suit: HEART, Value: 10},
	Card{Rank: QUEEN, Suit: DIAMOND, Value: 10},
	Card{Rank: QUEEN, Suit: SPADE, Value: 10},
	Card{Rank: QUEEN, Suit: CLUB, Value: 10},
	Card{Rank: KING, Suit: CLUB, Value: 10},
}

var handFlush5 CardSet = CardSet{
	Card{Rank: TWO, Suit: HEART, Value: 2},
	Card{Rank: THREE, Suit: HEART, Value: 3},
	Card{Rank: FOUR, Suit: HEART, Value: 4},
	Card{Rank: FIVE, Suit: HEART, Value: 5},
	Card{Rank: JOKER, Suit: HEART, Value: 10},
}

var handFlush4 CardSet = CardSet{
	Card{Rank: TWO, Suit: HEART, Value: 2},
	Card{Rank: THREE, Suit: HEART, Value: 3},
	Card{Rank: FOUR, Suit: HEART, Value: 4},
	Card{Rank: FIVE, Suit: HEART, Value: 5},
	Card{Rank: JOKER, Suit: DIAMOND, Value: 10},
}

var hand2Pair CardSet = CardSet{
	Card{Rank: TWO, Suit: HEART, Value: 2},
	Card{Rank: TWO, Suit: DIAMOND, Value: 2},
	Card{Rank: FOUR, Suit: HEART, Value: 4},
	Card{Rank: QUEEN, Suit: HEART, Value: 10},
	Card{Rank: JOKER, Suit: DIAMOND, Value: 10},
}

var hand3Pair CardSet = CardSet{
	Card{Rank: TWO, Suit: HEART, Value: 2},
	Card{Rank: TWO, Suit: DIAMOND, Value: 2},
	Card{Rank: TWO, Suit: SPADE, Value: 4},
	Card{Rank: QUEEN, Suit: HEART, Value: 10},
	Card{Rank: JOKER, Suit: DIAMOND, Value: 10},
}

var hand4Pair CardSet = CardSet{
	Card{Rank: TWO, Suit: HEART, Value: 2},
	Card{Rank: TWO, Suit: DIAMOND, Value: 2},
	Card{Rank: TWO, Suit: SPADE, Value: 4},
	Card{Rank: TWO, Suit: CLUB, Value: 4},
	Card{Rank: JOKER, Suit: DIAMOND, Value: 10},
}

var hand2x2Pair CardSet = CardSet{
	Card{Rank: TWO, Suit: HEART, Value: 2},
	Card{Rank: TWO, Suit: DIAMOND, Value: 2},
	Card{Rank: THREE, Suit: SPADE, Value: 3},
	Card{Rank: THREE, Suit: CLUB, Value: 3},
	Card{Rank: JOKER, Suit: DIAMOND, Value: 10},
}

func TestOverallScoring(t *testing.T) {

	// Test 29 hard total score.
	score := hand29.GetTotalScoreWithNob(&hand29CutCard)
	expectedScore := 29
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test 29 hand total without knob
	score = hand29.GetTotalScore()
	expectedScore = 28
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test double run hand total
	score = handDoubleRun.GetTotalScore()
	expectedScore = 12
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test quad run 24 hand
	score = hand24.GetTotalScore()
	expectedScore = 24
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test triple run 24 hand
	score = handTriplerun.GetTotalScore()
	expectedScore = 15
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test non-nob hand.
	score = nobHand.GetTotalScoreWithNob(&handCutCard2)
	expectedScore = 6
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

}

func TestGetFifteenScore(t *testing.T) {

	// Test 29 hand 15 score
	score := hand29.GetFifteenScore()
	expectedScore := 16
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test entire deck 15 score
	deck := CreateDeck()
	score = deck.GetFifteenScore()
	expectedScore = 34528
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}
}

func TestFlushScore(t *testing.T) {

	// Test flush with 5 same suits in row
	score := handFlush5.GetFlushScore()
	expectedScore := 5
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test flush with 4 same suits in row
	score = handFlush4.GetFlushScore()
	expectedScore = 4
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}
}

func TestPairScore(t *testing.T) {

	// Test 2 pair
	score := hand2Pair.GetPairScore()
	expectedScore := 2
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test 3 pair
	score = hand3Pair.GetPairScore()
	expectedScore = 6
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test 4 pair
	score = hand4Pair.GetPairScore()
	expectedScore = 12
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test 2x2 pair
	score = hand2x2Pair.GetPairScore()
	expectedScore = 4
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test pair score of entire deck
	deck := CreateDeck()
	score = deck.GetPairScore()
	expectedScore = 156
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}
}
