package main

import (
	"fmt"
)

type Address struct {
	City, State string
}

type Person struct {
	Name    string
	Age     int
	Address // Embedded struct
}

func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	fmt.Println(p.Name)
	fmt.Println(p.City) // Accessing City directly through Person, thanks to embedding
}
