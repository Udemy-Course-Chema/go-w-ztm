package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	fmt.Println(rooms)

	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		fmt.Println(room)
		if room.cleaned {
			fmt.Println(room.name, "is Clean")
		} else {
			fmt.Println(room.name, "is Dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office", cleaned: true},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)
}
