package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)
//
// 	productNames[2] = "A Car"
//
// 	fmt.Println(productNames)
//
// 	fmt.Println(prices[2])
//
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99) // append don't modify but instead makes a copy and adds it to new
	fmt.Println(prices)                   // still the same
	fmt.Println(updatedPrices)            // prices plus new value but not modifying prices

	prices = append(prices, 5.99) // modify / update prices
	fmt.Println(prices)

	prices = prices[1:] // new slice ommiting first element or deleteing
	fmt.Println(prices)
}
