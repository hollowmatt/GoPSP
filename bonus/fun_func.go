package main

// takes in two vars, returns a value
func plus(a, b int) int {
	return a + b
}

// takes in 3 vars, returns 2 values
func math(operation string, a, b int) (int, string) {
	switch operation {
	case "add":
		return (a + b), operation
	case "subtract":
		return (a - b), operation
	case "multiply":
		return (a * b), operation
	case "divide":
		return 0, operation
	default:
		return (a + b), "add"
	}
}
