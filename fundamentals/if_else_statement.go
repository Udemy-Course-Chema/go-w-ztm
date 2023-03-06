package main

import "fmt"

func main() {
	if true {
		fmt.Printf("Hello\n")
	} else {
		fmt.Printf("No estamos")
	}

	// ########################################
	var age = 20
	if age >= 20 {
		fmt.Println("Eres mayor de edad")
	} else if age < 20 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Número inválido")

	}

	// ########################################
	var hasTicket = true
	var isValid = false
	if hasTicket && isValid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}

	// ########################################
	if rank := getUserRank(); rank == "admin" {
		fmt.Println("Es admin")

	}
}

func getUserRank() string {
	return "admin"
}
