package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// )

// func getAllInstructions() []int {
// 	data, err := ioutil.ReadFile("day-5-input.txt")
// 	if err != nil {
// 		panic("Read error!, Quitting")
// 	}

// 	inputs := strings.Split(string(data), ",")

// 	instructions := []int{}
// 	for i := 0; i < len(inputs); i++ {
// 		res, err := strconv.Atoi(inputs[i])
// 		if err != nil {
// 			panic("Invalid data format, Quitting")
// 		}
// 		instructions = append(instructions, res)
// 	}

// 	return instructions
// }

// func parseInstruction(instruction int) map[string]int {
// 	var modes map[string]int
// 	instr := strconv.Itoa(instruction)
// 	insLen := len(instr)

// 	switch insLen {
// 	case 1:
// 		modes = map[string]int{"m1": 0, "m2": 0, "m3": 0, "instruction": instruction}
// 	case 2:
// 		modes = map[string]int{"m1": 0, "m2": 0, "m3": 0, "instruction": instruction}
// 	case 3:
// 		in, _ := strconv.Atoi(instr[1:])
// 		i, _ := strconv.Atoi(string(instr[0]))
// 		modes = map[string]int{"m1": i, "m2": 0, "m3": 0, "instruction": in}
// 	case 4:
// 		in, _ := strconv.Atoi(string(instr[2:]))
// 		i1, _ := strconv.Atoi(string(instr[0]))
// 		i2, _ := strconv.Atoi(string(instr[1]))
// 		modes = map[string]int{"m1": i2, "m2": i1, "m3": 0, "instruction": in}
// 	case 5:
// 		in, _ := strconv.Atoi(string(instr[3:]))
// 		i1, _ := strconv.Atoi(string(instr[1]))
// 		i2, _ := strconv.Atoi(string(instr[2]))
// 		i3, _ := strconv.Atoi(string(instr[0]))
// 		modes = map[string]int{"m1": i2, "m2": i1, "m3": i3, "instruction": in}
// 	}

// 	return modes
// }

// func getParamCount(instruction int) int {
// 	switch instruction {
// 	case 1:
// 		fallthrough
// 	case 2:
// 		return 3
// 	case 3:
// 		fallthrough
// 	case 4:
// 		return 1
// 	case 99:
// 		return 0
// 	}
// 	return -1
// }

// func runOP(op int, p1 int, p2 int, out int, modes map[string]int, instructions []int) int {
// 	param1, param2, output := 0, 0, 0

// 	if modes["m1"] == 0 {
// 		param1 = instructions[p1]
// 	} else {
// 		param1 = p1
// 	}

// 	if modes["m2"] == 0 {
// 		param2 = instructions[p2]
// 	} else {
// 		param2 = p2
// 	}

// 	if op == 1 {
// 		output = param1 + param2
// 		instructions[out] = output
// 	} else if op == 2 {
// 		output = param1 * param2
// 		instructions[out] = output
// 	} else if op == 3 {
// 		output = 1
// 		instructions[out] = output
// 		//fmt.Println("3", output, instructions[out])
// 	} else if op == 4 {
// 		output = instructions[p1]
// 		if output != 0 {
// 			fmt.Println(output)
// 		}
// 	} else {
// 		panic("POSSIBLE ERROR")
// 	}

// 	return output
// }

// func getParams(instructions []int, instruction int, paramcount int) map[string]int {
// 	var p1, p2, out int
// 	switch paramcount {
// 	case 3:
// 		p1 = instructions[instruction+1]
// 		p2 = instructions[instruction+2]
// 		out = instructions[instruction+3]
// 		return map[string]int{"p1": p1, "p2": p2, "out": out}
// 	case 1:
// 		p1 = instructions[instruction+1]
// 		return map[string]int{"p1": p1, "out": p1}
// 	case 0:
// 		return nil
// 	}
// 	return nil
// }

// func runDiagnostics(input int) {
// 	var cursor, paramCount int = 0, 0
// 	var modes map[string]int
// 	instructions := getAllInstructions()

// 	for cursor <= len(instructions) {
// 		modes = parseInstruction(instructions[cursor])
// 		paramCount = getParamCount(modes["instruction"])

// 		if paramCount == 0 {
// 			break
// 		}

// 		paramset := getParams(instructions, cursor, paramCount)

// 		_ = runOP(
// 			modes["instruction"],
// 			paramset["p1"],
// 			paramset["p2"],
// 			paramset["out"],
// 			modes,
// 			instructions,
// 		)

// 		cursor += paramCount + 1
// 		//fmt.Println(output)
// 	}
// }

// func main() {
// 	runDiagnostics(1)
// }
