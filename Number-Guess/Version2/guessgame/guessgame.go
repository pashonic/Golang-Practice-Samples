package guessgame

type guessResult int

const (
	TOOHIGH guessResult = iota
	TOOLOW
	WIN
	LOSS
	GAMEOVER
)

type GuessGame struct {
	Guesses      uint64
	RandomNumber uint64
}

func (self *GuessGame) Guess(guess uint64) guessResult {

	if self.Guesses < 1 {
		return GAMEOVER
	}

	if guess == self.RandomNumber {
		return WIN
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
