package main

import "fmt"

func main() {
	var value = hello()

	fmt.Println(value)

	var totalSuma = sum(5, 19)

	fmt.Println(totalSuma)

	doblar := double(2)

	fmt.Println(doblar)
}

func hello() string {
	fmt.Println("Helloooo!")

	return "perame"
}

func sum(a int, b int) int {
	var total = (a + b)
	return total
}

func double(x int) int {
	return x + x
}
