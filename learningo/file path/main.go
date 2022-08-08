package main

import (
	"fmt"
	"path"
)

func main() {

	var file string

	_, file = path.Split("./basics/main.go")

	fmt.Println("file :", file)

	//	Type Conversion
	speed := 100
	force := 2.5
	var sXf float64

	sXf = float64(speed) * force

	fmt.Println(sXf)

}
