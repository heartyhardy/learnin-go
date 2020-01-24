package aoc15

import (
	"fmt"
	"strconv"
)

func genNextSequence(in []int, steps int) int {

	if steps <= 0 {
		//If you want to check the returning array turned to a string
		/*
			s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(in)), ""), "[]")
			fmt.Println(s)
		*/
		return len(in)
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

	partI := func(steps int) int {
		startsq := make([]int, len(start))
		for k, v := range start {
			vi, _ := strconv.Atoi(string(v))
			startsq[k] = vi
		}
		res := genNextSequence(startsq, steps)
		return res
	}

	partII := func(steps int) int {
		startsq := make([]int, len(start))
		for k, v := range start {
			vi, _ := strconv.Atoi(string(v))
			startsq[k] = vi
		}
		res := genNextSequence(startsq, steps)
		return res
	}

	run40 := partI(40)
	run50 := partII(50)

	fmt.Println("\n-- AoC 2015: Day 10: Elves Look, Elves Say --")
	fmt.Printf("\n After 4️⃣0️⃣  iterations: %21v \n After 5️⃣0️⃣  iterations: %21v\n", run40, run50)
	fmt.Println("\n-- DONE --")
	fmt.Println("")
}
