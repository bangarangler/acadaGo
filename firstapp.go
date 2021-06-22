package main

import "fmt"

func main() {
	// long declaration
	// var greetingText string
	// greetingText = "Hello World!"

	// shorter declaration with type
	// var greetingText string = "Hello World!"
	// even shorter and infered type
	// var greetingText = "Hello World!"

	// infered type go will figure it out usually use
	greetingText := "Hello World!"
	luckyNumber := 12
	unluckyNumber := luckyNumber + 5

	fmt.Println(greetingText)
	fmt.Println(luckyNumber)
	fmt.Println(unluckyNumber)

	unluckyNumber = 19
	fmt.Println(unluckyNumber)
}
