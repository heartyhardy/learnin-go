package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

const (
	or = iota
	and
	not
	rshift
	lshift
	eq
)

func readLines(filename string) (content []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	content = strings.Split(string(data), "\n")
	return
}

func mapFunctions(content ...string) (fn map[string][]string) {
	fn = make(map[string][]string)
	for _, line := range content {
		fields := strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})

		fops := make([]string, 5)
		switch len(fields) {
		case 2:
			fops[4] = fields[len(fields)-1]
			if _, err := strconv.Atoi(fields[0]); err != nil {
				fops[1] = fields[0]
			} else {
				fops[3] = fields[0]
			}
		case 3:
			fops[0], fops[1], fops[4] = fields[0], fields[1], fields[len(fields)-1]
		case 4:
			fops[0], fops[1], fops[2], fops[4] = fields[1], fields[0], fields[2], fields[len(fields)-1]
		}
		fn[fields[len(fields)-1]] = fops
	}
	return
}

//Run solution
func Run() {
	content := readLines("./inputs/day-07.txt")
	fn := mapFunctions(content...)
	for k, v := range fn {
		fmt.Println(k, v)
	}
}
