package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// declare a constant for the length of the solution
const solutionLength = 5

// declare a new error type "errInvalidWordLength"
var errInvalidWordLength = fmt.Errorf("Invalid guess, word doesn't have the proper number of charcters")

type Game struct {
	reader *bufio.Reader
}

func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}
	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")
	fmt.Printf("Enter a Guess: \n")
	guess := g.ask()
	fmt.Printf("Your guess is: %s\n", string(guess))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := []rune(string(playerInput))

		err = g.validateGuess(guess)

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s. \n", err.Error())
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("%s - expected %d, got %d", errInvalidWordLength, solutionLength, len(guess))
	}
	return nil
}
