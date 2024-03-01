package main

import (
	"fmt"
)

type Base struct{}

func (Base) Describe() {
	fmt.Println("Base")
}

type Derived struct {
	Base // Embedding Base into Derived
}

func (Derived) Describe() {
	fmt.Println("Derived")
}

func main() {
	d := Derived{}
	d.Describe()
}
