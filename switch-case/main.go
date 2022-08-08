package main

import "fmt"

func main() {

	i := 142
	j := 80
	switch {
	case i > 100:
		fmt.Printf("big ")
		fallthrough
	case i > 0:
		fmt.Printf("positive ")
		fallthrough
	default:
		fmt.Printf("number %d \n", i)

	}
	switch {
	case j > 100:
		fmt.Printf("big ")
		fallthrough
	case j > 0:
		fmt.Printf("positive ")
		fallthrough
	default:
		fmt.Printf("number %d", j)

	}
}
