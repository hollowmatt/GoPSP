package main

import (
	"fmt"
	"hollowmatt/gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	g, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start game: %s", err)
		return
	}
	g.Play()
}
