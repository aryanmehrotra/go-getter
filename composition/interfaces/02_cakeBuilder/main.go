package main

import "fmt"

func main() {
	cake := NewCake("Truffle", &Frosting{}, &Sprinkles{})
	cake.Deliver()

	fmt.Println("------------------------")

	cheeseCake := NewCake("CheeseCake", &Frosting{})
	cheeseCake.Deliver()

	fmt.Println("------------------------")

	IceCake := NewCake("IceCake", &Sprinkles{})
	IceCake.Deliver()
}
