package main

import "fmt"

func main() {
	data := MiEstructura{"World", 1, 2}

	fmt.Println(data.field)

	otroDato := MiEstructura{}

	otroDato.field = "Hello"

	fmt.Println("El otro dato :", otroDato)

}

type MiEstructura struct {
	field string
	a, b  int
}
