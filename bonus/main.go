package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bonus files")
	mapper()
	stringly()
	fmt.Println(plus(1, 3))
	a, b := math("add", 1, 2)
	fmt.Println(a, b)
	sum(1, 2, 3, 4, 5)
	fmt.Println("=========================")

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("=========================")

	fmt.Println(fact(7))

	fmt.Println("=========================")
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
	fmt.Println("=========================")

	fmt.Println("+++++++++++++++++++++++++++")
	fmt.Println(person{"John", 20})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(person{name: "Bob"})
	fmt.Println(newPerson("Catnip"))

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Rex",
		isGood: true,
	}
	fmt.Println(dog)

}
