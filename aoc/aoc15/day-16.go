package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

const (
	children    = "children"
	cats        = "cats"
	samoyeds    = "samoyeds"
	pomeranians = "pomeranians"
	akitas      = "akitas"
	vizslas     = "vizslas"
	goldfish    = "goldfish"
	trees       = "trees"
	cars        = "cars"
	perfumes    = "perfumes"
)

type auntSue struct {
	id    int
	prop  string
	count int
}

type auntsList map[int][]*auntSue

var tickerTape = map[string]int{
	akitas:      0,
	children:    3,
	cats:        7,
	cars:        2,
	goldfish:    5,
	pomeranians: 3,
	perfumes:    1,
	samoyeds:    2,
	trees:       3,
	vizslas:     0,
}

func readAuntSueList(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func makeAuntsList(sues []string) auntsList {
	aunts := make(auntsList)
	for _, sue := range sues {
		tape := strings.FieldsFunc(sue, func(c rune) bool {
			return !unicode.IsNumber(c) && !unicode.IsLetter(c)
		})

		index := 0
		for i, comp := range tape[1:] {
			_, err := strconv.Atoi(comp)
			if err != nil {
				sueID, _ := strconv.Atoi(tape[1])
				count, _ := strconv.Atoi(tape[i+2])
				if _, ok := aunts[sueID]; !ok {
					aunts[sueID] = make([]*auntSue, 3)
				}
				aunts[sueID][index] = &auntSue{id: sueID, prop: comp, count: count}
				index++
			}
		}
	}
	return aunts
}

func findMatch(aunts auntsList, isOutdated bool) int {
	lastMatch, matchID := 0, 0
	for i, props := range aunts {
		matchIndex := 0
		for _, prop := range props {
			if isOutdated {
				if tickerTape[prop.prop] == prop.count {
					matchIndex++
				}
				continue
			}
			switch prop.prop {
			case trees:
				fallthrough
			case cats:
				if tickerTape[prop.prop] < prop.count {
					matchIndex++
				}
			case pomeranians:
				fallthrough
			case goldfish:
				if tickerTape[prop.prop] > prop.count {
					matchIndex++
				}
			default:
				if tickerTape[prop.prop] == prop.count {
					matchIndex++
				}
			}
		}
		if lastMatch < matchIndex {
			lastMatch = matchIndex
			matchID = i
		}
	}
	return matchID
}

//Run Solution
func Run() {
	contents := readAuntSueList("./inputs/day-16.txt")
	aunts := makeAuntsList(contents)
	auntSue := findMatch(aunts, true)
	realAuntSue := findMatch(aunts, false)

	fmt.Printf("\n\t -- Day 16: Aunt Sue --\n")
	fmt.Printf("\n👩\tAunt Sue: %39v \n\n🙋‍♀️\tReal Aunt Sue: %34v\n", auntSue, realAuntSue)
	fmt.Println(`
	______
	|  _  \
	| | | |___  _ __   ___
	| | | / _ \| '_ \ / _ \
	| |/ / (_) | | | |  __/
	|___/ \___/|_| |_|\___|`)
	fmt.Println("")
}
