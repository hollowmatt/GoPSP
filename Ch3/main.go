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
	fmt.Println((bookworms))

	books := bookCount(bookworms)
	fmt.Println((books))

}
