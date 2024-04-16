package money_test

import (
	"hollowmatt/moneychange/money"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount   money.Amount
		to       money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.19 USD to EUR": {
			amount: money.Amount{},
			to:     money.Currency{},
			validate: func(t *testing.T, got money.Amount, err error) {
				expected := money.Amount{}
				if err != nil {
					t.Errorf("expected %v, got %v", expected, got)
				}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %v, got %v", expected, got)
				}
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
	}
}
