package guessgame

type guessResult int

const (
	TOOHIGH guessResult = iota
	TOOLOW
	WON
	LOSS
)

type GuessGame struct {
	Guesses      int
	MaxValue     int
	RandomNumber int
}

func (self *GuessGame) Guess(guess int) guessResult {
	if guess == self.RandomNumber {
		return WON
	}

	self.Guesses--

	if self.Guesses < 1 {
		return LOSS
	}

	if guess > self.RandomNumber {
		return TOOHIGH
	} else {
		return TOOLOW
	}
}
