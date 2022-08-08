package main

import "fmt"

func main() {
	/*

		type gram float64
		type ounce float64

		var g gram = 1000
		var o ounce

		o = ounce(g) * 0.035274

		fmt.Printf("%g grams is %.2f ounces \n", g, o)*/

	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var apple Kilogram = 10
	var salt Gram = 100
	var truck Ton = 1

	salt = Gram(apple)
	apple = Kilogram(truck)
	truck = Ton(salt)

	fmt.Printf("%d %d %d", salt, apple, truck)
}
