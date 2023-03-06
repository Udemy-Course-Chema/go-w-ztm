package main

import "fmt"

type Rectangle struct {
	a int
	b int
}

func area(a, b int) int {
	return a * b
}

func perimetro(a, b int) int {
	return 2 * (a + b)
}

func main() {
	miRectangulo := Rectangle{5, 2}

	areaValue := area(miRectangulo.a, miRectangulo.b)
	fmt.Println("El area de un rectangulo es: ", areaValue)

	perimetroValue := perimetro(miRectangulo.a, miRectangulo.b)
	fmt.Println("El perimetro de un rectangulo es: ", perimetroValue)

}
