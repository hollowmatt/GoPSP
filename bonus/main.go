package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bonus files")
	mapper()
	stringly()
	fmt.Println(plus(1, 3))
	a, b := math("add", 1, 2)
	fmt.Println(a, b)
	sum(1, 2, 3, 4, 5)
	fmt.Println("=========================")

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("=========================")

	fmt.Println(fact(7))
}
