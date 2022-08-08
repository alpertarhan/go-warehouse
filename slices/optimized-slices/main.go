package main

import (
	"fmt"
	"strings"
)

func main() {

	//functions from databases or etc.
	tasks := []string{"jump", "run", "read"}

	//creating slice from almost staticly by tasks slice
	upTasks := make([]string, 0, len(tasks))

	// adding variables to new slice
	for _, task := range tasks {

		upTasks := append(upTasks, strings.ToUpper(task))

		fmt.Printf("%v", upTasks)
	}

}
