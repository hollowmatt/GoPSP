package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
	}
	//

	common := findCommonBooks(bookworms)
	fmt.Println("Common books are:")
	displayBooks(common)

}
