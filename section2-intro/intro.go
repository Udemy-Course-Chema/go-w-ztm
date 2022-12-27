package main

import (
	"fmt"
)

// Format Package

func main() {
	var myName = "Chm0x"
	var herName string = "América"
	miNombre := "Chaos Chema"

	fmt.Println("My name is: ", myName)
	fmt.Println("Her name is: ", herName)
	fmt.Println("Nombre asignado: ", miNombre)

	var sum int
	fmt.Println(" El valor es: ", sum)

	part1, other := "Parte 1 ", 2
	fmt.Println("Parte 1: ", part1, " Otro: ", other)

	var (
		lessonName       = "Go Programming Course"
		numberOfStudents = 200213
	)

	fmt.Println("Curso: ", lessonName, " y estudiantes: ", numberOfStudents)

	word1, word2, _ := "Hello", "World", "!"

	fmt.Println(word1, word2)
}
