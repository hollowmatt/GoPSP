package main

import (
	"encoding/json"
	"os"
)

// A Bookworm contains a list of books on a person's (bookworm's) shelf
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book author and title pair
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// read file, return list
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	// if there is a filepath error, we'll catch it here
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var bookworms []Bookworm //store array of bookworms
	// if we get here, there is a file
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}
