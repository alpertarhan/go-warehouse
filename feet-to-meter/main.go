package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	feetInMeters = 0.3048
	feetInYards  = feetInMeters / 0.9144
)

func main() {

	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Printf("Error occured, %q value is not a number.\n", arg)
		return
	}
	meters := feetInMeters * feet
	yards := math.Round(feetInYards * feet)

	fmt.Printf("%g Feet equal %g Meters.\n", feet, meters)
	fmt.Printf("%g Feet equal %g Yards.\n", feet, yards)
}
