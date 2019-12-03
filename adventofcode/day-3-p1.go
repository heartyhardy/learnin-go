package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type tuple struct {
	a int
	b int
}

func readInput() (w1 []string, w2 []string) {
	data, err := ioutil.ReadFile("day-3-input.txt")
	if err != nil {
		os.Exit(1)
	}

	wires := strings.Split(string(data), "\n")
	w1 = strings.Split(string(wires[0]), ",")
	w2 = strings.Split(string(wires[1]), ",")

	return
}

func findPoints(directions []string) {

	dir := map[string]tuple{
		"L": tuple{-1, 0},
		"R": tuple{1, 0},
		"U": tuple{0, 1},
		"D": tuple{0, -1}
	}

	for _, dn := range directions {
		d := dn[:1]
		s := dn[1:]

	}
}

func main() {
	w1, _ := readInput()
	findPoints(w1)
}
