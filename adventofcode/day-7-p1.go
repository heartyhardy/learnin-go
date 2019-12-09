package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var input int = 0

// func setInput(in int) {
// 	input = in
// }

// func readRawIntCodes(filename string) []int {
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	intCodestr := strings.Split(string(data), ",")

// 	var intCodes []int
// 	for _, strcode := range intCodestr {
// 		intcode, _ := strconv.Atoi(strcode)
// 		intCodes = append(intCodes, intcode)
// 	}
// 	return intCodes
// }

// func getParamModes(instruction int, params map[string]int) {
// 	code := (instruction%10)*1 + (instruction/10%10)*10
// 	mode1 := (instruction / 100) % 10
// 	mode2 := (instruction / 1000) % 10
// 	mode3 := (instruction / 10000) % 10

// 	params["code"] = code
// 	params["mode1"] = mode1
// 	params["mode2"] = mode2
// 	params["mode3"] = mode3
// }

// func getParams(instruction int, cursor int, intCodes []int) (map[string]int, int) {
// 	params := map[string]int{}
// 	paramCount := 0

// 	getParamModes(instruction, params)

// 	switch params["code"] {
// 	case 1:
// 		fallthrough
// 	case 2:
// 		if params["mode1"] == 0 {
// 			params["in1"] = intCodes[intCodes[cursor+1]]
// 		} else if params["mode1"] == 1 {
// 			params["in1"] = intCodes[cursor+1]
// 		}

// 		if params["mode2"] == 0 {
// 			params["in2"] = intCodes[intCodes[cursor+2]]
// 		} else if params["mode2"] == 1 {
// 			params["in2"] = intCodes[cursor+2]
// 		}

// 		params["out"] = intCodes[cursor+3]
// 		paramCount = 3

// 	case 3:
// 		params["in1"] = input
// 		params["out"] = intCodes[cursor+1]
// 		paramCount = 1
// 	case 4:
// 		params["in1"] = intCodes[cursor+1]
// 		paramCount = 1
// 	case 5:
// 		fallthrough
// 	case 6:
// 		if params["mode1"] == 0 {
// 			params["in1"] = intCodes[intCodes[cursor+1]]
// 		} else if params["mode1"] == 1 {
// 			params["in1"] = intCodes[cursor+1]
// 		}

// 		if params["mode2"] == 0 {
// 			params["in2"] = intCodes[intCodes[cursor+2]]
// 		} else if params["mode2"] == 1 {
// 			params["in2"] = intCodes[cursor+2]
// 		}

// 		paramCount = 2

// 	case 7:
// 		fallthrough
// 	case 8:
// 		if params["mode1"] == 0 {
// 			params["in1"] = intCodes[intCodes[cursor+1]]
// 		} else if params["mode1"] == 1 {
// 			params["in1"] = intCodes[cursor+1]
// 		}

// 		if params["mode2"] == 0 {
// 			params["in2"] = intCodes[intCodes[cursor+2]]
// 		} else if params["mode2"] == 1 {
// 			params["in2"] = intCodes[cursor+2]
// 		}

// 		params["out"] = intCodes[cursor+3]
// 		paramCount = 3
// 	case 99:
// 		params["in1"] = 99
// 		paramCount = 0
// 	default:
// 		panic("INVALID INSTRUCTION!!")
// 	}
// 	return params, paramCount
// }

// func add(params map[string]int, intCodes []int) {
// 	intCodes[params["out"]] = params["in1"] + params["in2"]
// }

// func multiply(params map[string]int, intCodes []int) {
// 	intCodes[params["out"]] = params["in1"] * params["in2"]
// }

// func copy(params map[string]int, intCodes []int) {
// 	intCodes[params["out"]] = params["in1"]
// }

// func output(params map[string]int, intCodes []int) {
// 	if intCodes[params["in1"]] != 0 {
// 		fmt.Printf("\u001b[34;1m------ OUTPUT : %v ------\n", intCodes[params["in1"]])
// 	}
// }

// func jump(params map[string]int, intCodes []int) int {
// 	switch params["code"] {
// 	case 5:
// 		if params["in1"] != 0 {
// 			return params["in2"]
// 		}
// 	case 6:
// 		if params["in1"] == 0 {
// 			return params["in2"]
// 		}
// 	}
// 	return 0
// }

// func compare(params map[string]int, intCodes []int) {
// 	switch params["code"] {
// 	case 7:
// 		if params["in1"] < params["in2"] {
// 			intCodes[params["out"]] = 1
// 		} else {
// 			intCodes[params["out"]] = 0
// 		}
// 	case 8:
// 		if params["in1"] == params["in2"] {
// 			intCodes[params["out"]] = 1
// 		} else {
// 			intCodes[params["out"]] = 0
// 		}
// 	}
// }

// func performInstructions(code int, cursor int, intCodes []int) (skip int, isSeek bool) {

// 	params, paramCount := getParams(code, cursor, intCodes)

// 	skip = paramCount
// 	isSeek = false

// 	switch params["code"] {
// 	case 1:
// 		add(params, intCodes)
// 	case 2:
// 		multiply(params, intCodes)
// 	case 3:
// 		copy(params, intCodes)
// 	case 4:
// 		output(params, intCodes)
// 	case 5:
// 		fallthrough
// 	case 6:
// 		if validSkip := jump(params, intCodes); validSkip > 0 {
// 			skip = validSkip
// 			isSeek = true
// 		}
// 	case 7:
// 		fallthrough
// 	case 8:
// 		compare(params, intCodes)
// 	case 99:
// 	default:
// 		panic("INVALID INSTRUCTION!!")
// 	}
// 	return
// }

// func parseInstructions(intCodes []int) int {
// 	cursor, skip := 0, 0
// 	isSeek := false

// 	fmt.Printf("\n\n----- \u001b[33mINTCODE COMPUTER STARTED!\u001b[0m -----\n\n")

// 	for cursor < len(intCodes) {
// 		skip, isSeek = performInstructions(intCodes[cursor], cursor, intCodes)

// 		if skip == 0 {
// 			fmt.Print("\n----- \u001b[35;1mEND OF INSTRUCTIONS\u001b[0m ------\n")
// 			break
// 		}

// 		if isSeek {
// 			cursor = skip
// 		} else {
// 			cursor += skip + 1
// 		}
// 	}
// 	return intCodes[0]
// }

// func ampItUp(phaseSetting int) {
// 	intCodes := readRawIntCodes("day-7-input.txt")
// 	setInput(phaseSetting)
// 	_ = parseInstructions(intCodes)
// }

// //4
// func main() {
// 	ampItUp(14)
// 	// intCodes := readRawIntCodes("day-7-input.txt")
// 	// setInput(1)
// 	// response := parseInstructions(intCodes)
// 	// fmt.Printf("\n---- \u001b[33mPART I - RESPONSE (0 POS): \u001b[32;1m %v  ----\n\n", response)
// }
