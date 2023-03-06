package main

import "fmt"

func main() {

	// VARIABLE
	var sample struct {
		field string
		a, b  int
	}
	sample.field = "Hello world"
	sample.a = 10

	fmt.Println("sample: ", sample)

	// Por default
	otroSample := struct {
		field string
		a, b  int
	}{
		"Hello",
		1, 2,
	}

	fmt.Println("Otro Sample: ", otroSample)

}
