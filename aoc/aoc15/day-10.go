package aoc15

import (
	"fmt"
	"strconv"
)

func genNextSequence(in []int, steps int) []int {

	if steps <= 0 {
		//If you want to check the returning array turned to a string
		/*
			s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(in)), ""), "[]")
			fmt.Println(s)
		*/
		return in
	}
	res := make([]int, 0)
	for i := 0; i < len(in); i++ {
		n := 1
		for j := i; j < len(in)-1; j++ {
			if in[j] == in[j+1] {
				n++
				continue
			}
			break
		}
		v := in[i]
		res = append(res, n)
		res = append(res, v)
		i = i + n - 1
	}
	return genNextSequence(res, steps-1)
}

//Run Solution
func Run() {
	start := `1113122113`

	startsq := make([]int, len(start))
	for k, v := range start {
		vi, _ := strconv.Atoi(string(v))
		startsq[k] = vi
	}
	run40 := genNextSequence(startsq, 40)
	run50 := genNextSequence(run40, 10)

	fmt.Println("\n-- AoC 2015: Day 10: Elves Look, Elves Say --")
	fmt.Printf("\n After 4️⃣ 0️⃣  iterations: %21v \n After 5️⃣ 0️⃣  iterations: %21v\n", len(run40), len(run50))
	fmt.Println("\n-- DONE --")
	fmt.Println("")
}
