package money

import (
	"fmt"
	"strconv"
	"strings"
)

type Decimal struct {
	// subunits is the amount of subunits.  Multiply it by the precision to get the real value
	subunits int64
	// number of subunits in a unit, expressed as a power of 10
	precision byte
}

const (
	//ErrInvalidDecimal - returned if decimal is malformed
	ErrInvalidDecimal = Error("unable to convert the decimal")

	//ErrTooLarge - returned if quantity is too large
	ErrTooLarge = Error("quantity over 10^12 is too large")
)

// convert a string into decimal representation
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")
	const maxDecimal = 12 // max num digits in a thousand billion

	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	cents, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := byte(len(fracPart))

	return Decimal{subunits: cents, precision: precision}, nil
}
