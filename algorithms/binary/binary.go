package binary

import "fmt"

func BinarySearch(arr []int, target int) int {
	fmt.Println("Target is: ", target)
	fmt.Println("Array is: ", arr)
	fmt.Println("--------------------------------")

	// empty array case
	if len(arr) == 0 {
		return -1
	}

	// array of 1 value (but, target may not be in array)
	if len(arr) == 1 {
		if arr[0] == target {
			return 0
		} else {
			return -1
		}
	}

	// setup the low point and high point for searching in the array
	low := 0
	high := len(arr) - 1
	mid := (low + high) / 2
	steps := 0

	fmt.Println("low:", low, ", high:", high, ", mid:", mid, ", steps:", steps)

	// loop through
	for low <= high {
		steps++
		if arr[mid] == target {
			return steps
		}
		if arr[mid] > target {
			high = mid - 1
			mid = (low + high) / 2
		} else if arr[mid] < target {
			low = mid + 1
			mid = (low + high) / 2
		}
		fmt.Println("low:", low, ", high:", high, ", mid:", mid, ", steps:", steps)
	}
	// not found
	return -1
}
