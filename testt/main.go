package main

import "fmt"

func main() {

	a := 453
	b := 0
	for a > 0 {
		b++
		a /= 10
	}
	fmt.Println(b)
	c := []int{12, 34, 45, 56, 78, 90, 23}
	fmt.Println(c[3:6])
}
