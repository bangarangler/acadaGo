package main

import "fmt"

func main() {
	// new array contains three hobbies
	hobbies := [3]string{"dance", "basketball", "development"}
	// output array to terminal
	fmt.Println(hobbies)
	// array meta data (len and cap)
	fmt.Println(len(hobbies))
	fmt.Println(cap(hobbies))
	// first element
	fmt.Println(hobbies[0])
	// second and third element
	fmt.Println(hobbies[1:])

	// new array slice with only first two hobbies
	newHobbies := hobbies[:2]
	fmt.Println(newHobbies)

	// slice with 2 elements
	courseGoals := []string{"learn go", "become a master ;)"}
	// print it out
	fmt.Println(courseGoals)
	// change the second element
	courseGoals[1] = "become a master! :)"
	// print it out
	fmt.Println(courseGoals)

	// add a third element
	courseGoals = append(courseGoals, "third goal")
	// print it
	fmt.Println(courseGoals)

	// bonus create "Product" struct with title, id, price, and create a dynamic
	// list of products (at least 2 products). then add a third product to
	// existing list of products
	type Product struct {
		id    string
		title string
		price float64
	}

	productsList := []Product{
		{id: "1", title: "product one", price: 4.44},
		{id: "2", title: "product two", price: 5.99},
	}
	fmt.Println(productsList)

	productsList = append(productsList, Product{id: "3", title: "product three", price: 12.45})
	fmt.Println(productsList)

	// extra remove first item from array
	productsList = productsList[1:]
	fmt.Println(productsList)
}
