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
	guesses      uint64
	randomNumber uint64
}

func New(guesses uint64, randomNumber uint64) *GuessGame {
	return &GuessGame{guesses, randomNumber}
}

func (self *GuessGame) GetGuesses() uint64 {
	return self.guesses
}

func (self *GuessGame) Guess(guess uint64) guessResult {

	if self.guesses < 1 {
		return GAMEOVER
	}

	if guess == self.randomNumber {
		return WIN
	}

	self.guesses--

	if self.guesses < 1 {
		return LOSS
	}

	if guess > self.randomNumber {
		return TOOHIGH
	}
	return TOOLOW
}
