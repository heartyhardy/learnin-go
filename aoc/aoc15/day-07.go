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

type store struct {
	wires       map[string]wire
	gates       map[string]gate
	values      map[string]int
	connections map[string][]string
}

type wire struct {
	v      int
	origin string
	solved bool
}

type gate struct {
	v        int
	lin      string
	rin      string
	out      string
	op       int
	resolved bool
}

func readLines(filename string) (content []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	content = strings.Split(string(data), "\n")
	return
}

func parseLine(ds *store, line string) {
	fields := strings.FieldsFunc(line, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	w := wire{origin: line}
	fn := gate{}

	switch len(fields) {
	case 2:
		if v, err := strconv.Atoi(fields[0]); err == nil {
			w.v = v
			w.solved = true
			ds.values[fields[1]] = v
		} else {
			fn.lin = fields[0]
			fn.out = fields[1]
			fn.op = eq
		}
	case 3:
		fn.op = not
		fn.lin = fields[1]
		fn.out = fields[2]
	case 4:
		switch fields[1] {
		case "AND":
			fn.op = and
		case "OR":
			fn.op = or
		case "LSHIFT":
			fn.op = lshift
		case "RSHIFT":
			fn.op = rshift
		}
		fn.lin = fields[0]
		fn.rin = fields[2]
		fn.out = fields[3]
	}
	ds.wires[fields[len(fields)-1]] = w
	ds.gates[line] = fn
}

func parseAll(content []string) (ds *store) {
	ds = new(store)
	ds.wires = make(map[string]wire)
	ds.gates = make(map[string]gate)
	ds.values = make(map[string]int)

	for _, line := range content {
		parseLine(ds, line)
	}
	return
}

func mapAssociations(ds *store) *store {
	ds.connections = make(map[string][]string)

	for fk, fn := range ds.gates {
		if fn.lin != "" {
			cons := ds.connections[fn.lin]
			cons = append(cons, fk)
			ds.connections[fn.lin] = cons
		}
		if fn.rin != "" {
			cons := ds.connections[fn.rin]
			cons = append(cons, fk)
			ds.connections[fn.rin] = cons
		}
	}
	return ds
}

//Run solution
func Run() {
	content := readLines("./inputs/day-07.txt")
	ds := parseAll(content)
	mapAssociations(ds)
}
