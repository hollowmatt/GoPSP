package gordle

import "strings"

// hint describes the validity of a character in a guess.
type hint byte

// feedback is the array of hints for the length of the word (guess)
type feedback []hint

const (
	absentCharacter hint = iota
	wrongPostion
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️"
	case wrongPostion:
		return "🟨"
	case correctPosition:
		return "🟩"
	default:
		return "💩"
	}
}

func (f feedback) String() string {
	result := strings.Builder{}
	for _, h := range f {
		result.WriteString(h.String())
	}
	return result.String()
}

func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}
	return true
}
