package cardset

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
	Rank Rank
	Suit Suit
}

type CardInterface interface{}

func Deal(in []CardInterface, count int) (out []CardInterface) {
	for i, _ := range rand.Perm(len(in))[:count] {
		out = append(out, in[i])
		in = append(in[:i], in[i+1:]...)
	}
	return
}
