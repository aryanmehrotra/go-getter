package main

type Options interface {
	apply(Delivery) Delivery
}
