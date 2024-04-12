package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// declare a constant for the length of the solution
const solutionLength = 5

// declare a new error type "errInvalidWordLength"
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the proper number of charcters")

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
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s. Expected %d, got %d \n", err.Error(), solutionLength, len(guess))
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return errInvalidWordLength
	}
	return nil
}

// split string into individual characters
func splitToUpperCaseChars(input string) []rune {
	return []rune(strings.ToUpper(input))
}
