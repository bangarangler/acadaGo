package main

import "fmt"

func main() {
	firstName := "Jon"
	var lastName string = "P"

	fmt.Println(firstName, lastName)

	currentYear := 2021
	var birthYear int
	birthYear = 1988

	age := currentYear - birthYear
	fmt.Println(age)

	currentYear = currentYear + 1
	age = currentYear - birthYear
	fmt.Println(age)

	// var (
	// 	firstName       = "Jon"
	// 	lastName        = "P"
	// 	currentYear int = 2021
	// 	birthYear   int = 1988
	// )
	//
	// firstName = "Jon"
	// lastName = "P"
	// age := currentYear - birthYear
	// ageNextYear := (currentYear + 1 - birthYear)
	//
	// fmt.Println(firstName + " " + lastName)
	// fmt.Println(firstName + " " + lastName)
	// fmt.Println("Your age is ", age)
	// fmt.Println("Your age next year is ", ageNextYear)
}
