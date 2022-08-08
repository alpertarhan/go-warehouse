package main

import (
	"fmt"
	"os"
)

func main() {

	inp := os.Args[1:]

	if len(inp) < 1 {
		fmt.Printf("You should give a website to mask it!")
		return
	}

	var (
		text      = inp[0]
		size      = len(text)
		stringBuf = make([]byte, 0, size)
		in        bool
	)
	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	for i := 0; i < size; i++ {

		stringBuf = append(stringBuf, text[i])

		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			stringBuf = append(stringBuf, link...)
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}
		stringBuf = append(stringBuf, c)
	}
	fmt.Println(string(stringBuf))
}
