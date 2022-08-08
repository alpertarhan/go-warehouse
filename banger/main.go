package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	arg := os.Args[1]
	len := utf8.RuneCountInString(arg)
	repeat := strings.Repeat("!", len)
	msg := strings.ToUpper(repeat + arg + repeat)

	fmt.Printf("%s", msg)

}
