package aoc15

import (
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

type gates map[string][]string

func readLines(filename string) (content []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	content = strings.Split(string(data), "\n")
	return
}

func mapFunctions(content ...string) (fn gates) {
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

func resolve(fn gates) {
	values := make(map[string]uint16)
	for k, v := range fn {
		// fmt.Println(k, v)
		// fmt.Println(sv, isVal, path)
		// fmt.Println(fn[path[len(path)-1]])
		path := make([]string, 0)
		fv, isSolved := canSolve(values, v)
		if isSolved {
			values[v[4]] = fv
			continue
		} else {

		}
		break
	}
}

func canSolve(values map[string]uint16, fnops []string) (val uint16, solved bool) {

	if len(fnops[0]) == 0 {
		vfn, err := strconv.ParseUint(fnops[3], 10, 16)
		if err != nil {
			if v, ok := values[fnops[1]]; !ok {
				return 0, false
			} else {
				val = v
				solved = true
				return
			}
		}
		val = uint16(vfn)
		solved = true
		return
	}

	vl, lerr := strconv.ParseUint(fnops[1], 10, 16)
	vr, rerr := strconv.ParseUint(fnops[2], 10, 16)
	var l, r uint16

	if lerr != nil {
		if v, ok := values[fnops[1]]; ok {
			l = v
		} else {
			return 0, false
		}
	} else {
		l = uint16(vl)
	}
	if rerr != nil {
		if v, ok := values[fnops[1]]; ok {
			r = v
		} else {
			return 0, false
		}
	} else {
		r = uint16(vr)
	}

	switch fnops[0] {
	case "NOT":
		val = ^l
	case "AND":
		val = l & r
	case "OR":
		val = l | r
	}
	solved = true
	return
}

func seek(fn string, fns gates, path *[]string) (v uint16) {
	ops := fns[fn]

	switch ops[0] {
	case "":
		//Save path
		//Save value
		//Return
	case "NOT":
		//Save path
		//Seek (L)
	case "AND":
		fallthrough
	case "OR":
		fallthrough
	case "LSHIFT":
		fallthrough
	case "RSHIFT":
		//Convert L  and R
		// If R is a Value and L is not -> Select L
		// If L is a Value and R is not -> Select R
		// If both are not Values, Select L
	}
}

//Run solution
func Run() {
	content := readLines("./inputs/day-07.txt")
	fn := mapFunctions(content...)
	resolve(fn)

}
