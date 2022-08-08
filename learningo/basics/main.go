package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//	print statements in Go
	println("Hello from my first go program")
	fmt.Println("Hello from my first go program with fmt package")
	fmt.Println("Number of my CPU :", runtime.NumCPU())

	//	decleration basics in Go
	var speed int
	var heat float64
	var off bool
	var brand string
	const pi = 3.14159

	//	short decleration -it can figure out its type by itself-
	myName := "Alper"
	myAge := 24
	velocity, isExpandable := 50, true
	//	to redecleration a variable with short decleration, you need to decleare at least one new variable
	velocity, isExpandable = 40, false
	velocity, isDecleratable := 35, true

	//	This is for better readability
	var (
		video string

		duration int
		current  int
	)

	var (
		force float64
		now   time.Time
	)
	force, now = 5.67, time.Now()

	fmt.Println(
		"\n		speed as  int :", speed,
		"\n		heat as float64 :", heat,
		"\n		off as bool :", off,
		"\n		brand as string :", brand,
		"\n		pi as float64 :", pi,
		"\n		Name :", myName,
		"\n		Age :", myAge,
		"\n\n	Velocity :", velocity,
		"\n 	Expandable :", isExpandable,
		"\n\n	Redeclaration :", isDecleratable,
		"\n\n	Video :", video,
		"\n\n	Duration :", duration,
		"\n\n	Current :", current,
		"\n\n	Force :", force,
		"\n\n	Now :", now)

}
