package aoc15

import (
	"fmt"
	"io/ioutil"
)

const (
	up    = byte('^')
	down  = byte('v')
	left  = byte('<')
	right = byte('>')
)

const (
	santa = iota
	bot
)

type xy struct {
	x int
	y int
}

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	return data
}

func deliverPresents(data []byte) int {
	x, y := 0, 0
	grid := make(map[xy]int)

	origin := xy{x, y}
	grid[origin] = 1

	for _, v := range data {

		switch v {
		case up:
			y++
		case down:
			y--
		case left:
			x--
		case right:
			x++
		}
		newXY := xy{x, y}
		grid[newXY]++
	}
	return len(grid)
}

func deliverPresentsDuo(data []byte) int {
	x1, y1, x2, y2 := 0, 0, 0, 0
	grid := make(map[xy]int)
	origin := xy{x1, y1}
	grid[origin] = 1
	turn := santa

	for _, v := range data {

		switch turn {
		case 0:
			switch v {
			case up:
				y1++
			case down:
				y1--
			case left:
				x1--
			case right:
				x1++
			}
			newXY := xy{x1, y1}
			grid[newXY]++
		case 1:
			switch v {
			case up:
				y2++
			case down:
				y2--
			case left:
				x2--
			case right:
				x2++
			}
			newXY := xy{x2, y2}
			grid[newXY]++
		}

		turn ^= 1
	}
	return len(grid)
}

//Run the solution
func Run() {
	data := readFile("./inputs/day-03.txt")
	santaRun := deliverPresents(data)
	santaDuo := deliverPresentsDuo(data)

	fmt.Println("\n-- AoC 2015: Day 3: Perfectly Spherical Houses in a Vacuum --")
	fmt.Printf("\nSolo Run: %v \nDuo Run with SantaBot: %v\n", santaRun, santaDuo)
	fmt.Println("\n-- DONE --")
	fmt.Println("")
}
