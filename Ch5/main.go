package main

import (
	"hollowmatt/gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {
	solution := "shake"
	g := gordle.New(os.Stdin, solution, maxAttempts)
	g.Play()
}
