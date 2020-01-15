package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	off = iota
	on
	toggle
)

const (
	onOff = iota
	brightness
)

func readAllLines(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	contents := strings.Split(string(data), "\n")
	return contents
}

func parseLine(line string) (opcode, cs, rs, ce, re int) {
	ops := strings.Fields(line)
	switch ops[0] {
	case "turn":
		if ops[1] == "on" {
			opcode = 1
		} else {
			opcode = 0
		}
		start := strings.Split(ops[2], ",")
		rs, _ = strconv.Atoi(start[0])
		cs, _ = strconv.Atoi(start[1])
		end := strings.Split(ops[4], ",")
		re, _ = strconv.Atoi(end[0])
		ce, _ = strconv.Atoi(end[1])
	case "toggle":
		opcode = 2
		start := strings.Split(ops[1], ",")
		rs, _ = strconv.Atoi(start[0])
		cs, _ = strconv.Atoi(start[1])
		end := strings.Split(ops[3], ",")
		re, _ = strconv.Atoi(end[0])
		ce, _ = strconv.Atoi(end[1])
	}
	return
}

func letThereBeLight(lines []string, pattern int) (on, off int) {
	var lights interface{} = setupLights(pattern)
	for _, line := range lines {
		op, cs, rs, ce, re := parseLine(line)
		switch pattern {
		case onOff:
			setLights(lights.([][]bool), onOff, op, cs, rs, ce, re)
		case brightness:
			setLights(lights.([][]int), brightness, op, cs, rs, ce, re)
		}
	}
	switch pattern {
	case onOff:
		return countEmUp(lights.([][]bool))
	case brightness:
		on = calcBrightness(lights.([][]int))
		off = 0
		return
	default:
		panic("Invalid mode!")
	}
}

func setLights(lights interface{}, pattern, op, cs, rs, ce, re int) {
	for r := rs; r <= re; r++ {
		for c := cs; c <= ce; c++ {
			switch pattern {
			case onOff:
				switch op {
				case on:
					lights.([][]bool)[r][c] = true
				case off:
					lights.([][]bool)[r][c] = false
				case toggle:
					lights.([][]bool)[r][c] = !lights.([][]bool)[r][c]
				}
			case brightness:
				switch op {
				case on:
					lights.([][]int)[r][c]++
				case off:
					lights.([][]int)[r][c]--
					if lights.([][]int)[r][c] < 0 {
						lights.([][]int)[r][c] = 0
					}
				case toggle:
					lights.([][]int)[r][c] += 2
				}
			}
		}
	}
}

func setupLights(pattern int) (lights interface{}) {
	switch pattern {
	case onOff:
		lights = make([][]bool, 1000)
		for r := range lights.([][]bool) {
			lights.([][]bool)[r] = make([]bool, 1000)
		}
	case brightness:
		lights = make([][]int, 1000)
		for r := range lights.([][]int) {
			lights.([][]int)[r] = make([]int, 1000)
		}
	}
	return
}

// PART I - Counting lights which are ON
func countEmUp(lights [][]bool) (on, off int) {
	for _, rows := range lights {
		for _, col := range rows {
			if col == true {
				on++
			} else {
				off++
			}
		}
	}
	return
}

// Part II - Calculating Total Brightness
func calcBrightness(lights [][]int) (brightness int) {
	for _, rows := range lights {
		for _, col := range rows {
			brightness += col
		}
	}
	return
}

func lightEmUp(lights [][]bool) {
	for _, rows := range lights {
		for _, col := range rows {
			if col == true {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

//Run solution
func Run() {
	content := readInstructions("./inputs/day-06.txt")
	on, off := letThereBeLight(content, onOff)
	totalBrightness, _ := letThereBeLight(content, brightness)

	fmt.Println("\n-- AoC 2015: Day 6: Probably a Fire Hazard --")
	fmt.Printf("\n💡  Turned (ON/OFF): %v / %v \n🔆  Total Brightness: %v", on, off, totalBrightness)
	fmt.Println("\n-- DONE --")
	fmt.Println("")
}
