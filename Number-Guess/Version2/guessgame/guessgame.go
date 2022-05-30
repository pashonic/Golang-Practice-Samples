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
	lockGame     bool
}

func (self *GuessGame) Guess(guess uint64) guessResult {

	if self.lockGame {
		return GAMEOVER
	}

	if guess == self.RandomNumber {
		return WIN
	}

	self.Guesses--

	if self.Guesses < 1 {
		self.lockGame = true
		return LOSS
	}

	if guess > self.RandomNumber {
		return TOOHIGH
	} else {
		return TOOLOW
	}
}
