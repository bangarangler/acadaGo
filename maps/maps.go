package main

import "fmt"

func main() {
	// works but have to know the order and if your not familiar with say aws not
	// clear
	// websites := []string{"https://google.com", "https://aws.com"}

	// create map
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	// print map
	fmt.Println(websites)
	// pring item in a map by key name
	fmt.Println(websites["Google"])
	fmt.Println(websites["Amazon Web Services"])

	// add new item to map
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites["LinkedIn"])

	// delete item from map
	delete(websites, "Google")
	fmt.Println(websites)

	// update item in map
	websites["LinkedIn"] = "https://linkedin.com/test"
	fmt.Println(websites["LinkedIn"])

}

// maps solve different problem than structs -> struct is set in stone and
// can't delete item from struct.  struct great for clearly defined data.

// if you don't have a structured set of data but want to label it a map could
// be a better choice as you can add to it and remove from it more strictly.

// type Websites struct {
//   google string // would have to know the names ahead of time
//   aws string
//   Amazon Web Services string // couldn't even do this
// }
// can think of maps as arrays but instead of indexes use any label of your choice
