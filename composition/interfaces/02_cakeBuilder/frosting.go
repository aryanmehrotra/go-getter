package main

import "fmt"

type Frosting struct{}

type FrostedCake struct {
	delivery   Delivery
	decoration string
}

func (f *Frosting) apply(delivery Delivery) Delivery {
	return &FrostedCake{delivery: delivery, decoration: "Adding frosts"}
}

func (d *FrostedCake) Deliver() {
	fmt.Println(d.decoration)
	d.delivery.Deliver()
}
