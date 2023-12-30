package recursion

import "fmt"

// This is not an Algorithm, but will be used by algorithms
func Recurse(counter int) {
	steps := 0
	fmt.Println("--------------------------------")
	fmt.Println("Recursion")
	fmt.Println("Countdown from: ", counter)
	fmt.Println("--------------------------------")
	countdown(counter, &steps)
	fmt.Println("Took ", steps, " steps")

}

func countdown(counter int, steps *int) {
	*steps++
	fmt.Print(counter, ", ")
	if counter < 1 {
		fmt.Print(" -> BLAST OFF")
		fmt.Println()
		return
	} else {
		countdown(counter-1, steps)
	}
}

func Plots(width int, length int) int {
	var big, little int
	size := 0

	if width == length {
		size = width
		fmt.Println("Lenght already equals width, no operations required")
		return size
	}

	if width > length {
		big = width
		little = length
	} else {
		big = length
		little = width
	}

	divide(big, little, &size)
	return size
}

func divide(big int, little int, size *int) {
	if big%little == 0 {
		*size = little
		return
	} else {
		newLittle := big % little
		newBig := little
		divide(newBig, newLittle, size)
	}
}
