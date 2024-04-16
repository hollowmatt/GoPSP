package money

type Decimal struct {
	// subunits is the amount of subunits.  Multiply it by the precision to get the real value
	subunits int64
	// number of subunits in a unit, expressed as a power of 10
	precision byte
}

ParseDecimal(string) (Decimal, error) {
	// 1 - find the position of the . and split on it. 
	// 2 - convert the string without the . to an integer. This could fail
	// 3 - add some consistency check
	// 4 - return the result
}
