package main

import (
	"testing"
)

var (
	jetBlackHeart = Book{Author: "Robert Gailbraith", Title: "Jet Black Heart"}
	paperWhite    = Book{Author: "Robert Gailbraith", Title: "Paper White"}
	soul          = Book{Author: "Tracy Kidder", Title: "Soul of the new Machine"}
	traffic       = Book{Author: "Tom Vanderbilt", Title: "Traffic: Why we drive the way we do"}
	billy         = Book{Author: "Stephen King", Title: "Billy Summers"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}

	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Sheila", Books: []Book{jetBlackHeart, paperWhite, traffic, billy}},
				{Name: "Matt", Books: []Book{soul, jetBlackHeart, traffic, paperWhite}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/nofile.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/badfile.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected err, got nothing")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			//compare contentss against expected
			if !equalBookworms(got, tc.want) {
				t.Fatalf("results do not match, got: %v, expected %v", got, tc.want)
			}
		})
	}
}

// compare two lists of bookworms, return true if they match
func equalBookworms(bookworms, target []Bookworm) bool {
	//compare length
	if len(bookworms) != len(target) {
		return false
	}

	//iterate and compare
	for i := range bookworms {
		// chack index name match
		if bookworms[i].Name != target[i].Name {
			return false
		}
		//check book list for the user
		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}
	return true
}

// compare two lists of books to see if they match
func equalBooks(books, target []Book) bool {
	//compare length
	if len(books) != len(target) {
		return false
	}

	//iterate and compare
	for i := range books {
		if (books[i].Author != target[i].Author) && (books[i].Title != target[i].Title) {
			return false
		}
	}

	return true
}

// Testing the book comparison
func TestBooksCount(t *testing.T) {
	testCase := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"base case (normal)": {
			input: []Bookworm{
				{Name: "Sheila", Books: []Book{jetBlackHeart, paperWhite, traffic}},
				{Name: "Matt", Books: []Book{soul, jetBlackHeart, traffic}},
			},
			want: map[Book]uint{jetBlackHeart: 2, paperWhite: 1, traffic: 2, soul: 1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm with no books": {
			input: []Bookworm{
				{Name: "Sheila", Books: []Book{jetBlackHeart, paperWhite, traffic}},
				{Name: "Matt", Books: []Book{}},
			},
			want: map[Book]uint{jetBlackHeart: 1, paperWhite: 1, traffic: 1},
		},
	}

	for name, tc := range testCase {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

// compare two book counts, see if they match
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()
	// if lengths don't match, it is an immediate fail
	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

func TestFindCommonBooks(t *testing.T) {
	testCase := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common book": {
			input: []Bookworm{
				{Name: "Sheila", Books: []Book{jetBlackHeart, paperWhite, traffic}},
				{Name: "Matt", Books: []Book{soul}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Sheila", Books: []Book{jetBlackHeart, paperWhite, traffic}},
				{Name: "Matt", Books: []Book{soul, jetBlackHeart}},
			},
			want: []Book{jetBlackHeart},
		},
	}

	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(test.input)
			if !equalBooks(test.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, test.want)
			}
		})
	}
}
