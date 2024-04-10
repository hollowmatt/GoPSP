package main

import (
	"hollowmatt/gordle/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}
