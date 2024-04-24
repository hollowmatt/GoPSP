package money

import (
	"fmt"
	"math"
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

	// a thousand billion (see above error)
	maxDecimal = 1e12
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

	dec := Decimal{subunits: cents, precision: precision}
	dec.simplify()

	return dec, nil
}

func (d *Decimal) simplify() {
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}

func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}

func (d *Decimal) String() string {
	if d.precision == 0 {
		return fmt.Sprintf("%d", d.subunits)
	}

	centsPerUnit := pow10(d.precision)
	frac := d.subunits % centsPerUnit
	integer := d.subunits / centsPerUnit

	decimalFormat := "%d.%0" + strconv.Itoa(int(d.precision)) + "d"
	return fmt.Sprintf(decimalFormat, integer, frac)
}
