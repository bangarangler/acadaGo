package main

import (
	"fmt"

	"github.com/bangarangler/acadaGo/greeting"
)

func main() {
	// long declaration
	// var greetingText string
	// greetingText = "Hello World!"

	// shorter declaration with type
	// var greetingText string = "Hello World!"
	// even shorter and infered type
	// var greetingText = "Hello World!"

	// infered type go will figure it out usually use
	// greetingText := "Hello World!"
	luckyNumber := 12
	unluckyNumber := luckyNumber + 5

	fmt.Println(greetingText)
	fmt.Println(greeting.GreetingTextFromGreeting)
	fmt.Println(luckyNumber)
	fmt.Println(unluckyNumber)

	unluckyNumber = 19
	fmt.Println(unluckyNumber)

	// newNumber := luckyNumber / 3 // 5.xxx --> 5
	var newNumber float64 = float64(luckyNumber) / 5 // 17 -> 17.0000
	fmt.Println(newNumber)

	var defaultFloat float64 = 1.236789123456789123456789
	var smallFloat float32 = 1.236789123456789123456789
	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	// "" for strings '' for special letters or emojies
	var firstRune rune = '💚' //128154
	fmt.Println(firstRune)
	fmt.Println(string(firstRune)) // prints the green heart

	var firstByte byte = 'a'
	fmt.Println(firstByte)         // 97
	fmt.Println(string(firstByte)) // a

	firstName := "Jon"
	lastName := "Palacio"
	// fullName := firstName + " " + lastName
	// fullName := fmt.Sprintln(firstName, " ", lastName)
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 33
	// fmt.Println("9" + 1) // nope
	// fmt.Println("Hi, my full name is " + fullName + " and I am " + age + "years old")
	fmt.Printf("Hi, I am %v and I am %v (Type: %T) years old.", fullName, age, age)
}
