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
}
