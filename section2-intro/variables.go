package main

import "fmt"

func main() {
	var color = "Black"
	fmt.Println(color)
	year, age := 1996, 26
	fmt.Println(year, age)
	var (
		first_initial = "J"
		last_initial  = "a"
	)
	fmt.Println(first_initial, last_initial)

	var age2 int
	age2 = 365 * age
	fmt.Println("Age times 365 days: ", age2)

}
