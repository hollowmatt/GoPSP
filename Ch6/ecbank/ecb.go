package ecbank

import "hollowmatt/moneychange/money"

// EuroCentralBank can call the bank to retrieve exchange rates.
type EuroCentralBank struct {
}

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
func (ecb EuroCentralBank) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	return 0, nil
}
