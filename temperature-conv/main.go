package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	celsius, _ := strconv.ParseFloat(arg, 64)

	fahrenheit := (celsius * 1.8) + 32

	fmt.Printf("%g Celcius equal %g Fahrenheit.", celsius, fahrenheit)

}
