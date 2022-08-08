package main

import "fmt"

func main() {

	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	fmt.Printf("Planet : %s\n", planet)
	fmt.Printf("Distance : %d million kms\n", distance)
	fmt.Printf("Orbital Period : %.3f days\n", orbital)
	fmt.Printf("Does %s have life on it? : %t\n", planet, hasLife)
	fmt.Printf("%s is %d km away. Think! %[2]d kms! %[1]s OMG!\n", planet, distance)
}
