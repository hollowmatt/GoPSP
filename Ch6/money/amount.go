package money

type Amount struct {
	quantity Decimal
	currency Currency
}

const ErrTooPrecise = Error("quantity is too precice")

// Return an amount of money
// NewAmount returns an Amount of money.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	switch {
	case quantity.precision > currency.precision:
		// In order to avoid converting 0.00001 cent, let's exit now.
		return Amount{}, ErrTooPrecise
	case quantity.precision < currency.precision:
		quantity.subunits *= pow10(currency.precision - quantity.precision)
		quantity.precision = currency.precision
	}
	return Amount{quantity: quantity, currency: currency}, nil

}

func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}
	return nil
}

func (a Amount) String() string {
	return a.quantity.String() + " " + a.currency.code
}
