package main

import "fmt"

// Sprinkles is a decorator that adds sprinkles to the cake.
type Sprinkles struct{}

type SprinkledCake struct {
	delivery   Delivery
	decoration string
}

func (s *Sprinkles) apply(delivery Delivery) Delivery {
	return &SprinkledCake{delivery: delivery, decoration: "Adding sprinkles"}
}

func (d *SprinkledCake) Deliver() {
	fmt.Println(d.decoration)
	d.delivery.Deliver()
}
