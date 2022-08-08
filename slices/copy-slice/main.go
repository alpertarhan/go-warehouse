package main

import "fmt"

func main() {

	odds := []string{"1", "3", "5"}
	even := []string{"0", "2", "4"}

	fmt.Printf("odds : %v\n even :%v", odds, even)

	numberOfCopied := copy(even, odds)
	fmt.Printf("\nodds : %v\n even :%v\n copied : %v", odds, even, numberOfCopied)
}
