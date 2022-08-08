package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	inp, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("%5s", "X")
	for i := 0; i <= inp; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Printf("\n")

	for i := 0; i <= inp; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= inp; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Printf("\n")
	}
}
