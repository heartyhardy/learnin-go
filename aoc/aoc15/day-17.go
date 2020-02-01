package aoc15

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readContainerData(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func findAllCombinations(containerData []string) int {
	containers := make([]int, len(containerData))
	for i, container := range containerData {
		icapacity, _ := strconv.Atoi(container)
		containers[i] = icapacity
	}
	sort.Ints(containers)

	//return recCount(containers, 150, len(containers)-1)
	return recCountMin(containers, 150, len(containers), 0)
}

func recCount(containers []int, target, i int) int {
	if target == 0 {
		return 1
	} else if target < 0 {
		return 0
	} else if i < 0 {
		return 0
	} else if target < containers[i] {
		return recCount(containers, target, i-1)
	} else {
		return recCount(containers, target-containers[i], i-1) +
			recCount(containers, target, i-1)
	}
}

func findMinCombinations(containerData []string) int {
	containers := make([]int, len(containerData))
	for i, container := range containerData {
		icapacity, _ := strconv.Atoi(container)
		containers[i] = icapacity
	}
	sort.Ints(containers)

	return recMin(containers, len(containers), 150)
}

func recCountMin(containers []int, target, size, i int) int {
	if size < 0 {
		return 0
	} else if target == 0 {
		return 1
	} else if i == len(containers) || target < 0 {
		return 0
	} else {
		return recCountMin(containers, target, size, i+1) +
			recCountMin(containers, target-containers[i], size-1, i+1)
	}
}

func recMin(containers []int, size, target int) int {
	i, min := 1, -1
	for min <= 0 {
		i++
		min = recCountMin(containers, 150, i, 0)
	}
	return min
}

//Run Solution
func Run() {
	containerData := readContainerData("./inputs/day-17.txt")
	maxCombinations := findAllCombinations(containerData)
	minCombinations := findMinCombinations(containerData)

	fmt.Printf("\n\t -- Day 17: No Such Thing as Too Much --\n")
	fmt.Printf("\n✳\tMax Combinations: %39v \n\n✳\tMin Combinations: %38v\n", maxCombinations, minCombinations)
	fmt.Println(`
	______
	|  _  \
	| | | |___  _ __   ___
	| | | / _ \| '_ \ / _ \
	| |/ / (_) | | | |  __/
	|___/ \___/|_| |_|\___|`)
	fmt.Println("")
}
