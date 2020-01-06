package aoc15

import (
	"fmt"
	"io/ioutil"
)

const (
	up   = byte('(')
	down = byte(')')
)

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	return data
}

func getFloors() (int, int) {
	cnt, pos := 0, 0
	input := readFile("./inputs/day-01.txt")

	for i, c := range input {
		switch c {
		case up:
			cnt++
		case down:
			cnt--
		}
		if cnt == -1 && pos == 0 {
			pos = i + 1
		}
	}
	return cnt, pos
}

// Run the solution
func Run() {
	cnt, pos := getFloors()
	fmt.Println("\n-- AoC 2015: Day 01 - Not Quite Lisp --")
	fmt.Printf("\nDestination Floor: %v \nFirst Basement Occurrence: %v", cnt, pos)
	fmt.Println("\n-- DONE --")
	fmt.Println("")
}
