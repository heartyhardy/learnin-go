package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type weightTable map[string]map[string]int

func readWeightTable(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	replaced := strings.ReplaceAll(string(data), "happiness units by sitting next to", "")
	replaced = strings.ReplaceAll(replaced, ".", "")
	contents := strings.Split(strings.ReplaceAll(replaced, "would", ""), "\n")
	return contents
}

func buildWeightTable(contents []string) weightTable {
	weights := make(weightTable)
	for _, line := range contents {
		fields := strings.Fields(line)

		if _, ok := weights[fields[0]]; !ok {
			weights[fields[0]] = make(map[string]int)
		}
		v, _ := strconv.Atoi(fields[2])
		switch fields[1] {
		case "lose":
			v *= -1
		}
		weights[fields[0]][fields[3]] = v
	}
	return weights
}

func getKeys(weights weightTable) []string {
	keys, i := make([]string, len(weights)), 0
	for k := range weights {
		keys[i], i = k, i+1
	}
	return keys
}

func arrange(weights weightTable, includesYou bool) int {
	keys := getKeys(weights)
	p := make([]int, len(keys))
	max := 0

	nextP := func() {
		for i := len(p) - 1; i >= 0; i-- {
			if i == 0 || p[i] < len(p)-i-1 {
				p[i]++
				return
			}
			p[i] = 0
		}
	}
	perm := func() []string {
		res := make([]string, len(keys))
		copy(res, keys)
		for k, v := range p {
			res[k], res[k+v] = res[k+v], res[k]
		}
		return res
	}
	calcHappiest := func(perm []string) int {
		sum := 0
		for i := 0; i < len(perm)-1; i++ {
			sum += weights[perm[i]][perm[i+1]]
			sum += weights[perm[i+1]][perm[i]]
		}
		if !includesYou {
			sum += weights[perm[0]][perm[len(perm)-1]]
			sum += weights[perm[len(perm)-1]][perm[0]]
		}
		return sum
	}

	for ; p[0] < len(p); nextP() {
		sum := calcHappiest(perm())
		if max < sum {
			max = sum
		}
	}

	return max
}

//Run Solution
func Run() {
	contents := readWeightTable("./inputs/day-13.txt")
	weights := buildWeightTable(contents)
	maxHappiness := arrange(weights, false)
	maxHappinessWithMeh := arrange(weights, true)

	fmt.Println("\n-- Day 13: Knights of the Dinner Table --")
	fmt.Printf("\n😀\tMax happiness: %35v \n😑\tMax Happiness after You join: %20v\n", maxHappiness, maxHappinessWithMeh)
	fmt.Println(`
	______
	|  _  \
	| | | |___  _ __   ___
	| | | / _ \| '_ \ / _ \
	| |/ / (_) | | | |  __/
	|___/ \___/|_| |_|\___|`)
	fmt.Println("")
}
