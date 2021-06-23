package main

import "fmt"

func main() {
	// tells go hey this is alaways the value. if I try to change it yell at me.
	const unchangingPI = 3.14
	// unchangingPI = 3.45 // nope nope nope ; )
	pi := 3.14
	radius := 5

	// circumference := 2 * int(pi) * radius // 2 * 3 * 5
	circumference := 2 * pi * float64(radius)

	fmt.Println(circumference)
	fmt.Printf("For a radius of %v, the circle circumference is %.2f", radius, circumference)

	// var PI float32 = 3.14
	// var circleRadius float32 = 5
	//
	// circleCircumference := (2 * PI * circleRadius)
	//
	// fmt.Println(circleCircumference)
	//
	// test := fmt.Sprintf("For a radius of %v, the circle circumference is %.2f", circleRadius, circleCircumference)
	//
	// fmt.Println(test)
}
