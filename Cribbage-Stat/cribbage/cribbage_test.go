package cribbage

import (
	"testing"
)

var hand29CutCard = Card{Rank: JOKER, Suit: HEART, Value: 10}
var hand29 CardSet = CardSet{
	Card{Rank: FIVE, Suit: HEART, Value: 5},
	Card{Rank: FIVE, Suit: DIAMOND, Value: 5},
	Card{Rank: FIVE, Suit: SPADE, Value: 5},
	Card{Rank: FIVE, Suit: CLUB, Value: 5},
	hand29CutCard,
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

}

func TestGetFifteenScore(t *testing.T) {

	// Test 29 hand 15 score.
	score := hand29.GetFifteenScore()
	expectedScore := 16
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

	// Test entire deck 15 score.
	deck := CreateDeck()
	score = deck.GetFifteenScore()
	expectedScore = 34528
	if score != expectedScore {
		t.Fatalf(`%v is not %v`, score, expectedScore)
	}

}
