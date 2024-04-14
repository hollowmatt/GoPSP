package gordle_test

import (
	"hollowmatt/gordle/gordle"
	"testing"
)

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file   string
		length int
		err    error
	}{
		"English dictionary": {
			file:   "../corpus/english.txt",
			length: 2315,
			err:    nil,
		},
		"Empty dictionsary": {
			file:   "../corpus/empty.txt",
			length: 0,
			err:    gordle.ErrCorpusEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := gordle.ReadCorpus(tc.file)
			if tc.err != err {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("expect %d, got %d", tc.length, len(words))
			}
		})
	}
}
