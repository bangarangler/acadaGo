package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// struct is like a object in javascript or a dictionary in python

type User struct { // uppercase if to be used in other packages lowercase if internal
	firstName string
	lastName  string
	birthdate string
	created   time.Time // custom type defined by go (struct from go team)
}

func (u *User) outputDetails() { // method (now coupled with the User struct)
	fmt.Printf("My name is %v %v (born on %v)", u.firstName, u.lastName, u.birthdate)
}

// func NewUser(fName, lName, bDate string) User {
// 	created := time.Now()
// 	user := User{  // thix is sub optimal will create and store 2 users in memory
// 		firstName: fName,
// 		lastName:  lName,
// 		birthdate: bDate,
// 		created:   created,
// 	}
// 	return user
// }

func NewUser(fName, lName, bDate string) *User {
	created := time.Now()
	user := User{
		firstName: fName,
		lastName:  lName,
		birthdate: bDate,
		created:   created,
	}
	return &user
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	// var newUser User
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYY): ")
	// created := time.Now()

	// newUser = User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// 	created:   created,
	// }

	// newUser = User{ // can omit beg if the order is always the same
	//   firstName,
	//   lastName,
	//   birthdate,
	//   created,
	// }

	// newUser = User{  // lastName will be created to default which is a "" for string
	// 	firstName: firstName,
	// 	birthdate: birthdate,
	// 	created:   created,
	// }
	newUser := NewUser(firstName, lastName, birthdate)

	// fmt.Println(firstName, lastName, birthdate, created)
	// fmt.Println(newUser)
	// outputUserDetails(newUser)
	newUser.outputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}
