package money

import "fmt"

// represents the rate to convert from one currency to another
type ExchangeRate Decimal

func Convert(amount Amount, to Currency) (Amount, error) {
	// Convert to the target currency applying the fetched change rate.
	convertedValue := applyExchangeRate(amount, to, ExchangeRate{subunits: 2, precision: 0})

	// validate the converted amount is in the handled bounded range.
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

func applyExchangeRate(amnt Amount, target Currency, rate ExchangeRate) (Amount, error) {
	converted, err := multiply(a.quantity, rate)
	if err != nil {
		return Amount{}, err
	}

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}, nil
}

func multiply(d Decimal, r ExchangeRate) (Decimal, error) {
	rate, err := ParseDecimal(fmt.Sprintf("%g", r))
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: exchange rate is %f", ErrInvalidDecimal, r)
	}
	dec := Decimal{
		subunits:  d.subunits * rate.subunits,
		precision: d.precision + rate.precision,
	}

	dec.simplify()
	return dec, nil
}
