package main

import "fmt"

func main() {
	greet("Jose")

	msg := message()

	fmt.Println("Mensaje desde la función: ", msg)

	total := add(53, 32, 23)

	fmt.Println("El total es: ", total)

	anyNum := anyNumber(10)
	fmt.Println("Cualquier numero: ", anyNum)
}

func greet(person string) {
	fmt.Println("Hello, ", person, " !!! How are you?")
}
func message() string {
	return "World"
}
func add(a, b, c int) int {
	return a + b + c
}
func anyNumber(a int) int {
	return a
}
