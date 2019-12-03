// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// )

// func getstrings() []int {
// 	data, err := ioutil.ReadFile("day-2-input.txt")
// 	if err != nil {
// 		print("Reading data file failed. Cannot skip this criical operation.")
// 		panic("Quitting")
// 	}

// 	inputs := strings.Split(string(data), ",")

// 	instructions := []int{}
// 	for i := 0; i < len(inputs); i++ {
// 		res, err := strconv.Atoi(inputs[i])
// 		if err != nil {
// 			print("Invalid data format!")
// 			panic("Quitting")
// 		}
// 		instructions = append(instructions, res)
// 	}

// 	return instructions
// }

// func doAdd(in1, in2, out1 int, ops []int) {
// 	res := ops[in1] + ops[in2]
// 	ops[out1] = res
// }

// func doMult(in1, in2, out1 int, ops []int) {
// 	res := ops[in1] * ops[in2]
// 	ops[out1] = res
// }

// func execInstructions(noun, verb int) int {

// 	l, i, j, k := 0, 0, 0, 0
// 	op, in1, in2, out1 := 0, 0, 0, 0
// 	ops := getstrings()

// 	ops[1] = noun
// 	ops[2] = verb

// 	for o := 0; o < len(ops)-4; o = o + 4 {

// 		l = o
// 		i = o + 1
// 		j = o + 2
// 		k = o + 3

// 		op = ops[l]
// 		in1 = ops[i]
// 		in2 = ops[j]
// 		out1 = ops[k]

// 		switch op {
// 		case 1:
// 			doAdd(in1, in2, out1, ops)
// 		case 2:
// 			doMult(in1, in2, out1, ops)
// 		case 99:
// 			break
// 		}
// 	}
// 	return ops[0]
// }

// func main() {
// 	res := 0
// 	for i := 0; i < 100; i++ {
// 		for j := 0; j < 100; j++ {
// 			res = execInstructions(i, j)

// 			if res == 19690720 {
// 				fmt.Println(100*i, j)
// 			}
// 		}
// 	}
// }
