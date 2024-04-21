package main

import (
	"flag"
	"fmt"
	"hollowmatt/moneychange/money"
	"os"
)

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "GBP", "target currency")

	// Pull out the currencies from flags
	flag.Parse()
	// grab the amount from args
	value := flag.Arg(0)

	//validate there is a value
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	fromCurr, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse the source currency %q: %s \n", *from, err.Error())
		os.Exit(1)
	}

	toCurr, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse the destination currency %q: %s \n", *to, err.Error())
		os.Exit(1)
	}

	quantity, err := money.ParseDecimal(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s. \n", value, err.Error())
		os.Exit(1)
	}

	amount, err := money.NewAmount(quantity, fromCurr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "shit fucked up %s \n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Amount: ", amount, ", Currency: ", toCurr)
}
