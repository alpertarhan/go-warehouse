package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		input := os.Args[i]
		fmt.Printf("%[1]d. input : %[2]v\n", i, input)
	}

}
