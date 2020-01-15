package aoc15

import (
	"io/ioutil"
	"strings"
)

const (
	rulesA = iota
	rulesB
)

func readInstructions(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	content := strings.Split(string(data), "\n")
	return content
}

func naughtyOrNice(content []string, ruleSet int) int {
	nc := 0
	for _, line := range content {
		switch ruleSet {
		case rulesA:
			if notContain(line) &&
				hasRepeatingChars(line) &&
				containsThreeVs(line) {
				nc++
			}
		case rulesB:
			if pairAppearsTwice(line) &&
				repeatsAfterOne(line) {
				nc++
			}
		}
	}
	return nc
}

//Rule 3 - does not contain the strings ab, cd, pq, or xy
func notContain(line string) bool {
	comps := []string{"ab", "cd", "pq", "xy"}
	for _, v := range comps {
		if strings.Contains(line, v) {
			return false
		}
	}
	return true
}

//Rule 2 - contains at least one letter that appears twice in a row
func hasRepeatingChars(line string) bool {
	lc := rune(line[0])
	for i, c := range line {
		if i == 0 {
			continue
		}
		if lc == c {
			return true
		}
		lc = c
	}
	return false
}

//Rule 1 - contains at least three vowels
func containsThreeVs(line string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	cnt := 0
	for _, v := range vowels {
		scnt := strings.Count(line, v)
		cnt += scnt
		if cnt >= 3 {
			return true
		}
	}
	return false
}

//New Rule 1 - contains a pair of any two letters that appears at least twice (NO Overlapping)
func pairAppearsTwice(line string) bool {
	for i := 1; i < len(line); i++ {
		cpair := line[i-1 : i+1]
		if strings.Count(line, cpair) >= 2 {
			return true
		}
	}
	return false
}

//New Rule 2 - contains at least one letter which repeats with exactly one letter between them
func repeatsAfterOne(line string) bool {
	for i := 0; i < len(line); i++ {
		if (i + 2) < len(line) {
			if line[i] == line[i+2] {
				return true
			}
		}
	}
	return false
}

//Run Solution
// func Run() {
// 	content := readInstructions("./inputs/day-05.txt")
// 	niceRulesA := naughtyOrNice(content, rulesA)
// 	niceRulesB := naughtyOrNice(content, rulesB)

// 	fmt.Println("\n-- AoC 2015: Day 5: Doesn't He Have Intern-Elves For This? --")
// 	fmt.Printf("\n😇 Or 😈 (Ruleset A): %v \n😇 Or 😈 (Ruleset B): %v", niceRulesA, niceRulesB)
// 	fmt.Println("\n-- DONE --")
// 	fmt.Println("")

// }
