package gordle

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, _ := New(strings.NewReader(tc.input), []string{string(tc.want)}, 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want = %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("HELLO"),
			expected: nil,
		},
		"too short": {
			word:     []rune("HE"),
			expected: errInvalidWordLength,
		},
		"too long": {
			word:     []rune("Happen"),
			expected: errInvalidWordLength,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, _ := New(nil, []string{"SLICE"}, 0)
			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "shake",
			solution:         "shake",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter": {
			guess:            "small",
			solution:         "hello",
			expectedFeedback: feedback{absentCharacter, absentCharacter, absentCharacter, correctPosition, wrongPostion},
		},
		"right but wrong position": {
			guess:            "hlelo",
			solution:         "hello",
			expectedFeedback: feedback{correctPosition, wrongPostion, wrongPostion, correctPosition, correctPosition},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("guess: %q, got the wrong feedback, expected %v, got %v", tc.guess, tc.expectedFeedback, fb)
			}
		})
	}
}
