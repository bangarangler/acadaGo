package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	// userInput, _ := reader.ReadString('\n')
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Invalid input...")
		return
	}

	isOldEnough := userAge >= 30

	if isOldEnough && userAge < 50 || userAge >= 60 {
		fmt.Println("You're eligible for our senior jobs.")
	} else if userAge >= 50 {
		fmt.Println("The best age!")
	} else if userAge >= 18 {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("Sorry not old enough")
	}

	fmt.Println("Goodbye!")

	// switch
	// switch userInput {
	// case "1":
	// 	fmt.Println("Option 1")
	// case "2":
	// 	fmt.Println("Option 2")
	// case "3":
	// 	fmt.Println("Option 3")
	// default:
	// 	fmt.Println("invalid choice")
	// }
}
