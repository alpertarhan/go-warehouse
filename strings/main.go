package main

import "fmt"

func main() {

	start, stop := 'A', 'Z'

	for n := start; n <= stop; n++ {
		fmt.Printf("%c => %[1]d\n", n)
	}
}
