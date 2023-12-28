package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println("The sum of ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("is: ", total)
}
