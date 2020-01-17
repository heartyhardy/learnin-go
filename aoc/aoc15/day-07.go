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
		fn[fields[len(fields)-1]] = fields
	}
	return
}

func resolve(fn map[string][]string) {
	track := make([]string, 0)
	for _, v := range fn {
		var l, r []string
		switch len(v) {
		case 2:
			_, lerr := strconv.Atoi(v[0])
			if lerr != nil {
				l = seek(fn, &track, v[0])
			} else {
				l = v
			}
		case 3:
			l = seek(fn, &track, v[1])
		case 4:
			_, lerr := strconv.Atoi(v[0])
			_, rerr := strconv.Atoi(v[2])

			if lerr != nil {
				l = seek(fn, &track, v[0])
			} else if rerr != nil {
				r = seek(fn, &track, v[2])
			}
		}
		fmt.Println("LEFT: ", l)
		fmt.Println("RIGHT: ", r)
	}
}

func seek(fn map[string][]string, track *[]string, node string) []string {
	if v, ok := fn[node]; ok && len(v) == 2 {
		*track = append(*track, v[0])
		return *track
	} else if !ok {
		return nil
	}

	v := fn[node]
	switch len(v) {
	case 3:
		*track = append(*track, v[1])
		return seek(fn, track, v[1])
	case 4:
		_, lerr := strconv.Atoi(v[0])
		_, rerr := strconv.Atoi(v[2])

		if lerr != nil {
			*track = append(*track, v[0])
			return seek(fn, track, v[0])
		} else if rerr != nil {
			*track = append(*track, v[2])
			return seek(fn, track, v[2])
		}
	}
	return nil
}

//Run solution
func Run() {
	content := readLines("./inputs/day-07.txt")
	fn := mapFunctions(content...)
	resolve(fn)
}
