package recursion

import "fmt"

// This is not an Algorithm, but will be used by algorithms
func Recurse(counter int) {
	fmt.Println("--------------------------------")
	fmt.Println("Recursion")
	fmt.Println("Countdown from: ", counter)
	fmt.Println("--------------------------------")
	countdown(counter)

}

func countdown(counter int) {
	fmt.Print(counter, ", ")
	if counter < 1 {
		fmt.Print(" -> BLAST OFF")
		fmt.Println()
		return
	} else {
		countdown(counter - 1)
	}
}
