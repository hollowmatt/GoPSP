package selection

import "fmt"

// Big O for this is O(n2)
// n squared time due to double looping required
func SelectionSort(arr []int) ([]int, int) {
	fmt.Println("--------------------------------")
	fmt.Println("Selection Sort")
	fmt.Println("Input array is: ", arr)
	fmt.Println("--------------------------------")
	steps := 0

	// loop through the entire array
	for i, _ := range arr {
		minIndex := i
		// loop through looking for smallest in remainder of array
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			steps++
		}
		// swap the values, putting smallest in i, then next time through, i moves up...
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		steps++
	}
	// return the sorted array
	return arr, steps
}
