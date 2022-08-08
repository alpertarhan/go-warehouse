package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	moods := [...][3]string{
		{"happy", "awesome", "good"},
		{"sad", "terrible"},
	}

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf("[Your Name]")
		return
	}

	uName, mood := args[0], args[1]

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	var mi int
	if mood != "positive" {
		mi = 1
	}

	fmt.Printf("%s feels %s\n", uName, moods[mi][n])

}
