package main

import (
	"fmt"
)

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}

	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2, Boarded: true}
	)

	if bill.Boarded {
		fmt.Println("Permitido")

	}

	var ella Passenger

	ella.Name = "Ella"
	ella.Boarded = true
	ella.TicketNumber = 201

	fmt.Println("Variable ella", ella)

	camion := Bus{ella}
	fmt.Println(camion)
	fmt.Println(camion.FrontSeat.TicketNumber)

}
