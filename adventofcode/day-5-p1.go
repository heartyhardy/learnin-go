package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type middleware = func(params map[string]int, intCodes []int)

// var alarm1202 middleware = func(params map[string]int, intCodes []int) {
// 	intCodes[params["in1"]] = params["val1"]
// 	intCodes[params["in2"]] = params["val2"]

// 	fmt.Printf("\n\n----- \u001b[33mALARM 1202 STATE INITIATED!\u001b[0m -----\n\n")
// }

// func runMiddleware1201(mw middleware, intCodes []int) {
// 	params := map[string]int{"in1": 1, "in2": 2, "val1": 12, "val2": 2}
// 	mw(params, intCodes)
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
// 		params["in1"] = 1
// 		params["out"] = intCodes[cursor+1]
// 		params["type"] = 3
// 		paramCount = 1
// 	case 4:
// 		params["in1"] = intCodes[cursor+1]
// 		params["type"] = 4
// 		paramCount = 1
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

// func performInstructions(code int, cursor int, intCodes []int) (skip int) {

// 	params, paramCount := getParams(code, cursor, intCodes)
// 	skip = paramCount

// 	switch params["code"] {
// 	case 1:
// 		add(params, intCodes)
// 	case 2:
// 		multiply(params, intCodes)
// 	case 3:
// 		copy(params, intCodes)
// 	case 4:
// 		output(params, intCodes)
// 	case 99:
// 	default:
// 		panic("INVALID INSTRUCTION!!")
// 	}
// 	return
// }

// func parseInstructions(intCodes []int) int {
// 	cursor := 0
// 	skip := 0

// 	fmt.Printf("\n\n----- \u001b[33mINTCODE COMPUTER STARTED!\u001b[0m -----\n\n")

// 	for cursor < len(intCodes) {
// 		skip = performInstructions(intCodes[cursor], cursor, intCodes)

// 		if skip == 0 {
// 			fmt.Print("\n----- \u001b[35;1mEND OF INSTRUCTIONS\u001b[0m ------\n")
// 			break
// 		}

// 		cursor += skip + 1
// 	}
// 	return intCodes[0]
// }

// func main() {
// 	intCodes := readRawIntCodes("day-5-input.txt")
// 	// RUN THIS IF PART I
// 	//runMiddleware1201(alarm1202, intCodes)
// 	response := parseInstructions(intCodes)
// 	fmt.Printf("\n---- \u001b[33mPART I - RESPONSE (0 POS): \u001b[32;1m %v  ----\n\n", response)
// }
