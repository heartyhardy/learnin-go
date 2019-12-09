package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	positionMode  = "0"
	immediateMode = "1"
	relativeMode  = "2"
)

var arglens map[int]int = map[int]int{1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2}

func exec(src []int, args []int, cursor int) (flush []int, exitcode int) {
	exitcode = 1
	const bufferSize = 3000
	relativeBase, iptcursor := 0, 0
	codes := make([]int, bufferSize)
	flush = make([]int, 0)
	copy(codes, src)

	for {
		fmtcode := fmt.Sprintf("%05d", codes[cursor])
		opcode, _ := strconv.Atoi(fmtcode[3:5])

		if opcode == 99 {
			return flush, 0
		}

		modes := map[int]string{
			1: fmtcode[2:3],
			2: fmtcode[1:2],
			3: fmtcode[0:1],
		}

		peek := func(pos int) int {
			switch modes[pos] {
			case positionMode:
				return codes[codes[cursor+pos]]
			case immediateMode:
				return codes[cursor+pos]
			case relativeMode:
				return codes[relativeBase+codes[cursor+pos]]
			}
			panic("(@peek) Invalid mode: " + modes[pos])
		}

		save := func(pos int, val int) {
			switch modes[pos] {
			case positionMode:
				codes[codes[cursor+pos]] = val
				return
			case relativeMode:
				codes[relativeBase+codes[cursor+pos]] = val
				return
			}
			panic("(@save) invalid mode: " + modes[pos])
		}

		switch opcode {
		case 1:
			save(3, peek(1)+peek(2))
			cursor += arglens[opcode]
		case 2:
			save(3, peek(1)*peek(2))
			cursor += arglens[opcode]
		case 3:
			save(1, args[iptcursor])
			iptcursor++
			cursor += arglens[opcode]
		case 4:
			flush = append(flush, peek(1))
			cursor += arglens[opcode]
		case 5:
			if peek(1) != 0 {
				cursor = peek(2)
			} else {
				cursor += arglens[opcode]
			}
		case 6:
			if peek(1) == 0 {
				cursor = peek(2)
			} else {
				cursor += arglens[2]
			}
		case 7:
			if peek(1) < peek(2) {
				save(3, 1)
			} else {
				save(3, 0)
			}
			cursor += arglens[opcode]
		case 8:
			if peek(1) == peek(2) {
				save(3, 1)
			} else {
				save(3, 0)
			}
			cursor += arglens[opcode]
		case 9:
			relativeBase += peek(1)
			cursor += arglens[opcode]
		default:
			panic("Invalid opcode: " + string(opcode))
		}
	}
}

func parse(input string) (intCodes []int) {
	scodes := strings.Split(strings.ReplaceAll(input, "\r\n", ""), ",")
	intCodes = make([]int, len(scodes))
	for i, str := range scodes {
		intCodes[i], _ = strconv.Atoi(str)
	}
	return
}

func boostkey(input string) string {
	intCodes := parse(string(input))
	output, _ := exec(intCodes, []int{1}, 0)
	if len(output) != 1 {
		fmt.Printf("OUTPUT: %v/n", output)
	}
	return fmt.Sprintf("%s%s", "Key: ", strconv.Itoa(output[len(output)-1]))
}

func sos(input string) string {
	intCodes := parse(string(input))
	output, _ := exec(intCodes, []int{2}, 0)
	if len(output) != 1 {
		fmt.Printf("OUTPUT: %v/n", output)
	}
	return fmt.Sprintf("%s%s", "SOS: ", strconv.Itoa(output[len(output)-1]))
}

func main() {
	input, _ := ioutil.ReadFile("day-9-input.txt")
	keycode := boostkey(string(input))
	fmt.Println("\n\u001b[34m------------------- BOOST KEYCODE ---------------------\u001b[0m")
	fmt.Printf("\t\t \u001b[32m%v\u001b[0m                   \n", keycode)
	fmt.Println("\u001b[34m-------------------------------------------------------\u001b[0m")
	fmt.Println()

	soscoords := sos(string(input))
	fmt.Println("\n\u001b[34m------------------- SOS COORDINATES --------------------\u001b[0m")
	fmt.Printf("\t\t\t \u001b[31m%v\u001b[0m                   \n", soscoords)
	fmt.Println("\u001b[34m--------------------------------------------------------\u001b[0m")
	fmt.Println()
}
