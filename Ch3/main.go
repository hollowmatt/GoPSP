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
	fmt.Println(bookworms)

	books := booksCount(bookworms)
	fmt.Println(books)

	common := findCommonBooks(bookworms)
	for _, book := range common {
		fmt.Println("Author: " + book.Author + ", Title: " + book.Title)
	}

}
