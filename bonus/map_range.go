package main

import (
	"fmt"
)

func mapper() {
	recipe := map[string]int{"potatoes": 4, "onions": 1, "garlic": 2, "sweet potatoes": 1}

	for name, qty := range recipe {
		fmt.Println("You need ", qty, name)
	}
}
