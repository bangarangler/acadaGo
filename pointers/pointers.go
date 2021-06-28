package main

import "fmt"

func main() {
	age := 33

	fmt.Println(age)

	// var myAge *int  // creating manually
	// myAge = &age
	myAge := &age

	fmt.Println(*myAge)

	*myAge = 34

	fmt.Println(*myAge)
	fmt.Println(age)

	// doubledAge := double(age) // copied new entry in memory.
	doubledAge := double(myAge) // pass address and in function modify by working on pointer
	fmt.Println(doubledAge)
	fmt.Println(age)
}

func double(num *int) int {
	result := *num * 2
	*num = 100
	return result
}
