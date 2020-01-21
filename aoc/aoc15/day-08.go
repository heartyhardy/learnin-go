package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readLines(filename string) (content []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	content = strings.Split(string(data), "\n")
	return
}

// Hard way

func processEachLine(content []string) int {
	total := 0
	for _, line := range content {
		line = line[1:len(line)]
		ucodes := strings.Count(line, "\\x") * 4
		dslashs := strings.Count(line, "\\\\") * 2
		dquotes := strings.Count(line, "\"")

		total += ucodes + dslashs + dquotes
	}
	return total - 1
}

func encodeEachLine(content []string) int {
	total := 0
	for _, line := range content {
		ucodes := strings.Count(line, "\\") + 2
		dslash := strings.Count(line, string('"'))

		total += ucodes + dslash
	}
	return total
}

// Easy way : Thanks to Astrus

func unquote(str string) string {
	s, _ := strconv.Unquote(str)
	return s
}

func quote(str string) string {
	return strconv.Quote(str)
}

func countBeforeEncode(content []string) int {
	total := 0
	for _, s := range content {
		total += len(s) - len(unquote(s))
	}
	return total
}

func countAfterEncode(content []string) int {
	total := 0
	for _, s := range content {
		total += len(quote(s)) - len(s)
	}
	return total
}

//Run Solution
func Run() {
	content := readLines("./inputs/day-08.txt")
	// total := processEachLine(content)
	// encodedtotal := encodeEachLine(content)

	totalBefore := countBeforeEncode(content)
	totalAfter := countAfterEncode(content)

	fmt.Println("\n-- AoC 2015: Day 8: Matchsticks --")
	fmt.Printf("\n Extra characters: %v \n After Encoding: %v", totalBefore, totalAfter)
	fmt.Println("\n-- DONE --")
	fmt.Println("")
}
