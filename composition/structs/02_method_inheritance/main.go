package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

// Method belonging to Animal
func (a Animal) Speak() {
	fmt.Printf("%s says: Hello!\n", a.Name)
}

type Dog struct {
	Animal // Embedding Animal into Dog
	Breed  string
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Golden Retriever",
	}

	d.Speak() // Accessing Speak method, which is defined in Animal
}
