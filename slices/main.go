package main

import "fmt"

func main() {

	items := []string{"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"astreoids", "simcity", "metroid", "defender", "rayman", "tempest", "ultima"}

	const pageSize = 3

	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := items[from:to]
		header := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		fmt.Println(header, currentPage)
	}
}
