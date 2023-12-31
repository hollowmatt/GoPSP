package quicksort

import "fmt"

// Big O for this is O(___)
// ___ time based on number of entries in array
func Qsort(arr []int) ([]int, int) {
	steps := 0
	fmt.Println("--------------------------------")
	fmt.Println("Quicksort")
	fmt.Println("Input Array is: ", arr)
	fmt.Println("--------------------------------")

	if len(arr) < 2 {
		fmt.Println("array of 0 or 1 elemetns already sorted, ", steps, " steps")
		return arr, steps
	}
	return quick(arr, &steps), steps

}

func quick(arr []int, steps *int) []int {
	var less, more []int
	var pivot int
	mid := len(arr) / 2
	if mid == 0 {
		pivot = 0
	} else {
		pivot = arr[mid]
	}

	if len(arr) < 2 {
		return arr
	} else {
		for idx, val := range arr {
			*steps++
			if idx != mid {
				if val <= pivot {
					less = append(less, val)
				} else {
					more = append(more, val)
				}
			}
		}

		return append(append(quick(less, steps), pivot), quick(more, steps)...)
	}
}
