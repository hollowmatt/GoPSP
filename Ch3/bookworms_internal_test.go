package main

import (
	"fmt"
	"testing"
)

var (
	jetBlackHeart = Book{Author: "Robert Gailbraith", Title: "Jet Black Heart"}
	paperWhite    = Book{Author: "Robert Gailbraith", Title: "Paper White"}
	soul          = Book{Author: "Tracy Kidder", Title: "Soul of the new Machine"}
	traffic       = Book{Author: "Tom Vanderbilt", Title: "Traffic: Why we drive the way we do"}
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
				{Name: "Matt", Books: []Book{soul, jetBlackHeart, traffic}},
				{Name: "Sheila", Books: []Book{jetBlackHeart, paperWhite, traffic}},
			},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if tc.wantErr {
				if err == nil {
					t.Fatal("expected err, bot noghitn")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			//all good
			fmt.Println(got)
		})
	}
}
