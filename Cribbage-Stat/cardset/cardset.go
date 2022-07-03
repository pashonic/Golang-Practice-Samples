package cardset

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

func CreateDeck() (deck cardSet) {

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
