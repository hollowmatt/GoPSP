package main

// The bookworm/reader
type Reader struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// A given book
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// Recommendation with confidence score
type Recommendation struct {
	Book  Book
	Score float64
}

func recommend(allReaders []Reader, target Reader, n int) []Recommendation {

}
