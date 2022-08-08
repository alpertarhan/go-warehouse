package main

import (
	"fmt"
	"os"
)

const usage = `
			   Requires Username or Password!
			   -----------------------------				
			go run main.go [username] [password]
			
			`

func main() {

	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}
	username := os.Args[1]
	password := os.Args[2]

	if username == "Jack" && password == "1888" {
		fmt.Println("Access Granted to : ", username)
	} else if !(username == "Jack") || !(password == "1888") {
		fmt.Println("Access Denied for : ", username)

	}

}
