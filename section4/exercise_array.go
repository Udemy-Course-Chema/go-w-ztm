package main

import "fmt"

type Product struct {
	price int
	name  string
}

func printStats(list [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	} // End For

	fmt.Println("Last Item of list: ", list[totalItems-1])
	fmt.Println("Total Items: ", totalItems)
	fmt.Println("Cost: ", cost)
}

func main() {
	shoppingList := [4]Product{
		{1, "Banana"},
		{10, "Apple"},
		{6, "Tomatoes"},
	}

	printStats(shoppingList)

	shoppingList[3] = Product{8, "wines"}

	printStats(shoppingList)
}
