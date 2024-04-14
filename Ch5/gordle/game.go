package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// declare a new error type "errInvalidWordLength"
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the proper number of charcters")

// Game holds all the information needed to play Gordle
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUpperCaseChars(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb.String())
		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉You won!  You found it in %d guess(es)!  The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}
	fmt.Printf("💩You lose!  The solution was: %s.  Better luck next time!\n", string(g.solution))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUpperCaseChars(string(playerInput))

		err = g.validateGuess(guess)

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s. Expected %d, got %d \n", err.Error(), len(g.solution), len(guess))
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return errInvalidWordLength
	}
	return nil
}

// split string into individual characters
func splitToUpperCaseChars(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func computeFeedback(guess, solution []rune) feedback {
	//initialize holders
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Error: guess and solution are different lengths.  Guess must be %d characters", len(solution))
	}

	// check for correct letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	// check for correct letters in incorrect positions
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			// already been processed
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				// already been processed
				continue
			}

			if character == target {
				result[posInGuess] = wrongPostion
				used[posInSolution] = true
				break
			}

		}
	}
	return result
}
