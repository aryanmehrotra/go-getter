package main

import "fmt"

// Cake represents a basic cake.
type Cake struct {
	name string
}

func (c Cake) Deliver() {
	fmt.Println(c.name)
}

// NewCake creates a new cake and applies decorators to it.
func NewCake(name string, option ...Options) Delivery {
	cake := &Cake{name: name}

	var delivery Delivery

	delivery = cake

	for _, o := range option {
		delivery = o.apply(delivery)
	}

	return delivery
}
