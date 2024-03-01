package main

func main() {
	cake := NewCake("Truffle", &Frosting{}, &Sprinkles{})
	cake.Deliver()
}
