package main

import (
	"fmt"
	"time"

	"hollowmatt.com/algorithms/breadth"

	"hollowmatt.com/algorithms/binary"
	"hollowmatt.com/algorithms/quicksort"
	"hollowmatt.com/algorithms/recursion"
	"hollowmatt.com/algorithms/selection"
)

func main() {
	// Algorithm 1: Binary Search O(log n)
	ary := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 19
	fmt.Println("Operations: ", binary.BinarySearch(ary, target))
	fmt.Println("================================")

	// Algorithm 2: Selection Sort O(n2)
	//unsorted := []int{5, 7, 3, 2, 4, 1, 6, 8, 7, 9}
	unsorted := []int{33, 10, 15, 7, 5, 1, 17, 56, 98, 3, 88, 4, 18, 91, 12}
	results := []int{}
	var steps int
	results, steps = selection.SelectionSort(unsorted)
	fmt.Println("results: ", results, " in ", steps, " steps")
	fmt.Println("================================")

	// Chapter 3: An aside for recursion
	recursion.Recurse(10)
	fmt.Println("================================")

	// Ch 4 - Algorithm 3: Quicksort
	fmt.Println("largest equal square plots of 1680x640 is: ", recursion.Plots(1680, 640))
	qlist := []int{33, 10, 15, 7, 5, 1, 17, 56, 98, 3, 88, 4, 18, 91, 12}
	sorted := []int{}
	sorted, steps = quicksort.Qsort(qlist)
	fmt.Println("results: ", sorted, " in ", steps, " steps")
	fmt.Println("================================")

	//Ch 6 - Algorithm 4: Breadth First Search (with Graph)
	graph := map[int][]string{
		88: {"17th", "21st", "28th", "Negly", "Target", "Traders", "Bakery"},
		86: {"17th", "40th", "Bigelow", "Target", "Bakery"},
	}
	stops, route := breadth.Breadthfirst("Bakery", graph)
	if stops == 0 && route == 0 {
		fmt.Println("No route found")
	} else {
		fmt.Println("Least number of stops to Bakery Sq is ", stops, " on route ", route)
	}
	fmt.Println("================================")

	// concurrency

	fmt.Println("time is: ", time.Now())
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("now time is: ", time.Now())
}
