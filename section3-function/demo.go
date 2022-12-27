package main

import "fmt"

func main() {

	greet()

	numero := 5
	doblar := double(5)
	fmt.Println("El doble del numero: ", numero, ", es: ", doblar)

	prueba := add(54, 32)
	fmt.Println("Prueba de la función add : ", prueba)
}

func double(x int) int {
	return x + x
}

// Con el int al final, todo se convierte el mismo tipo
func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello, mr Wayne")
}
