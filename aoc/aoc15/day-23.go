package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInstructionsList(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	replaced := strings.ReplaceAll(string(data), ",", "")
	return strings.Split(replaced, "\n")
}

func parseInstructions(instructionsList []string, is_overriden bool) {
	registers := make(map[string]int)
	instructions := make([][]string, len(instructionsList))
	for i, instruction := range instructionsList {
		ops := strings.Fields(instruction)
		instructions[i] = make([]string, len(ops))
		instructions[i] = ops
	}

	if is_overriden {
		registers["a"] = 1
	}

	i := 0
	for {
		if i >= len(instructions) {
			break
		}
		switch instructions[i][0] {
		case "hlf":
			if _, ok := registers[instructions[i][1]]; !ok {
				registers[instructions[i][1]] = 0
			}
			registers[instructions[i][1]] /= 2
			i++
		case "tpl":
			if _, ok := registers[instructions[i][1]]; !ok {
				registers[instructions[i][1]] = 0
			}
			registers[instructions[i][1]] *= 3

			i++
		case "inc":
			if _, ok := registers[instructions[i][1]]; !ok {
				registers[instructions[i][1]] = 0
			}
			registers[instructions[i][1]]++
			i++
		case "jmp":
			v, _ := strconv.Atoi(instructions[i][1])
			i += v
		case "jie":
			if v, _ := registers[instructions[i][1]]; v%2 == 0 {
				jv, _ := strconv.Atoi(instructions[i][2])
				i += jv
				continue
			}
			i++
		case "jio":
			if v, _ := registers[instructions[i][1]]; v == 1 {
				jv, _ := strconv.Atoi(instructions[i][2])
				i += jv
				continue
			}
			i++
		default:
			break
		}
	}
	fmt.Println("A: ", registers["a"], "B: ", registers["b"])
}

//Run solution
func Run() {
	instructions := readInstructionsList("./inputs/day-23.txt")
	parseInstructions(instructions, false)
	parseInstructions(instructions, true)
}
