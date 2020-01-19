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

type gates map[string][]string
type links map[string][]string
type values map[string]uint16

func readLines(filename string) (content []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	content = strings.Split(string(data), "\n")
	return
}

func replaceAndReadLines(filename string, a string, b string) (content []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	newContent := strings.Replace(string(data), a, b, 1)
	content = strings.Split(newContent, "\n")
	return
}

func mapFunctions(content ...string) (fn gates, cns links, vns values) {
	fn = make(gates)
	cns = make(links)
	vns = make(values)
	for _, line := range content {
		fields := strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})

		fops := make([]string, 5)
		switch len(fields) {
		case 2:
			fops[4] = fields[len(fields)-1]

			if v, isNum := isNumeric(fields[0]); isNum {
				fops[3] = fields[0]
				vns[fops[4]] = v
			} else {
				fops[0], fops[1] = "EQ", fields[0]
				mapConnections(fops[4], cns, fops[4], fops[1])
			}
		case 3:
			fops[0], fops[1], fops[4] = fields[0], fields[1], fields[len(fields)-1]
			mapConnections(fops[4], cns, fops[4], fops[1])
		case 4:
			fops[0], fops[1], fops[2], fops[4] = fields[1], fields[0], fields[2], fields[len(fields)-1]
			mapConnections(fops[4], cns, fops[4], fops[1], fops[2])
		}
		fn[fields[len(fields)-1]] = fops
	}
	return
}

func mapConnections(w string, cns links, ops ...string) {
	for _, op := range ops {
		if _, isNum := isNumeric(op); !isNum && len(op) > 0 {
			if _, ok := cns[op]; !ok {
				cns[op] = make([]string, 0)
			}
			cns[op] = append(cns[op], w)
		}
	}
}

func trySolve(fns gates, cns links, vns values) {

	isSolved := func(w string) bool {
		if ops, ok := fns[w]; ok {
			switch ops[0] {
			case "":
				_, lok := tryFetchValue(ops[3], vns)
				return lok
			case "EQ":
				l, lok := tryFetchValue(ops[1], vns)
				if lok {
					vns[ops[4]] = l
					return lok
				}
			case "NOT":
				l, lok := tryFetchValue(ops[1], vns)
				if _, exists := vns[ops[4]]; !exists && lok {
					vns[ops[4]] = ^l
					return lok
				}
			default:
				l, lok := tryFetchValue(ops[1], vns)
				r, rok := tryFetchValue(ops[2], vns)
				if lok && rok {
					switch ops[0] {
					case "OR":
						vns[ops[4]] = l | r
					case "AND":
						vns[ops[4]] = l & r
					case "LSHIFT":
						vns[ops[4]] = l << r
					case "RSHIFT":
						vns[ops[4]] = l >> r
					}
					return true
				}
			}
		}
		return false
	}

	for len(vns) < len(fns) {
		for w := range vns {
			for _, l := range cns[w] {
				_ = isSolved(l)
			}
		}
	}
}

func isNumeric(w string) (uint16, bool) {
	v, err := strconv.ParseUint(w, 10, 16)
	if err != nil {
		return 0, false
	}
	return uint16(v), true
}

func tryFetchValue(w string, vns values) (val uint16, ok bool) {
	if v, isNum := isNumeric(w); isNum {
		return uint16(v), true
	} else if !isNum && len(w) > 0 {
		if vl, ok := vns[w]; ok {
			return vl, true
		}
		return 0, false
	}
	return 0, false
}

//Run solution
func Run() {
	content := readLines("./inputs/day-07.txt")
	fn, cns, vns := mapFunctions(content...)
	trySolve(fn, cns, vns)
	sigA, sigB := strconv.Itoa(int(vns["a"])), strconv.Itoa(int(vns["b"]))

	content = replaceAndReadLines("./inputs/day-07.txt", sigB, sigA)
	fn, cns, vns = mapFunctions(content...)
	trySolve(fn, cns, vns)

	fmt.Println("\n-- AoC 2015: Day 7: Some Assembly Required --")
	fmt.Printf("\n ⚡ :(Wire A): %v \n ⚡⚡:(Wire A Override): %v", sigA, vns["a"])
	fmt.Println("\n\n-- DONE --")
	fmt.Println("")
}
