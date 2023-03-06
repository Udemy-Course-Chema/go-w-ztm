package main

import "fmt"

func main() {

	// Classic Loop For
	for i := 0; i < 10; i++ {
		fmt.Println("En el loop", i)

	}

	// WHILE
	var i = 0
	for i < 10 {
		fmt.Println("Que pedo, este es el While?", i)
		i++
		if i == 5 {
			break
		}
	}
}
