package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	want := "Hello world"
	got := greet()

	if got != want {
		// mark test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
