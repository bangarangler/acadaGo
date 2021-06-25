package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := generateRandomNumbers()
	ans := add(a, b)
	fmt.Println(ans)
	printNumber(ans)
}

func generateRandomNumbers() (r1, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)
	return
}

func add(x, y int) (sum int) {
	sum = x + y
	return // auto return sum as we named it above
}

func printNumber(number int) {
	fmt.Printf("The number is: %v", number)
}
