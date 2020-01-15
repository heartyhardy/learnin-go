package aoc15

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	or = iota
	and
	not
	rshift
	lshift
)

type store struct {
	wires map[string]wire
	gates map[string]gate
}

type wire struct {
	v      int
	origin gate
	gates  map[string]gate
}

type gate struct {
	v      int
	lin    string
	rin    string
	out    int
	op     int
	solved bool
}

func readLines(filename string) (content []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	content = strings.Split(string(data), "\n")
	return
}

func parseFunctions(content []string) {
	for _, v := range content {

	}
}

//Run solution
func Run() {
	content := readLines("./inputs/day-07.txt")
	fmt.Println(content)
}
