package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting!")
		return
	}

	if selectedChoice == "1" {
		calculateSumUpToNumber()
	} else if selectedChoice == "2" {
		calculateFactorial()
	} else if selectedChoice == "3" {
		calculateSumManually()
	} else {
		calculateListSum()
	}
}

func calculateSumUpToNumber() {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}
	fmt.Println(chosenNumber)
	sum := 0

	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i
	}
	fmt.Printf("Result: %v", sum)
}

func calculateFactorial() { // 5 => 5 * 4 * 3 * 2 * 1 => 120
	fmt.Println("Please enter your number: ")
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	factorial := 1

	for i := 1; i <= chosenNumber; i++ {
		factorial = factorial * i
	}
	fmt.Printf("Result: %v", factorial)
}

func calculateSumManually() {
	// want to keep iterating as the user is entering numbers
	isEnteringNumbers := true
	sum := 0

	fmt.Println("Keep on entering numbers, program will quit when you enter any invalid number")

	// Go while loop or do this code while isEnteringNumbers is true
	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		sum = sum + chosenNumber

		// if err != nil { // can be shortened as below
		//   isEnteringNumbers = false
		// }
		isEnteringNumbers = err == nil
	}

	fmt.Printf("Result: %v\n", sum)
}

func calculateListSum() {
	fmt.Println("Please enter a comma-seperated list of numbers.")
	inputNumberList, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	inputNumberList = strings.Replace(inputNumberList, "\n", "", -1)
	// "1,2,3,4,10" <- desired input from user
	inputNumbers := strings.Split(inputNumberList, ",") // convert to slice

	sum := 0

	for idx, value := range inputNumbers {
		fmt.Printf("Index: %v, Value: %v\n", idx, value)
		number, err := strconv.ParseInt(value, 0, 64)

		if err != nil {
			fmt.Println("Invalid input must enter x,x,x,xx (format: number and , no spaces)")
			continue // skip if error
			// break // break out of loop
		}

		sum = sum + int(number)
	}

	fmt.Printf("Result: %v\n", sum)
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)
	if err != nil {
		return 0, err
	}

	return int(chosenNumber), nil
}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the numbers up to number X")
	fmt.Println("2) Built the factorial up to number X")
	fmt.Println("3) sum up manually entered numbers")
	fmt.Println("4) sum up list of entered numbers")

	fmt.Print("Please make your choice: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT!")
	}
}
