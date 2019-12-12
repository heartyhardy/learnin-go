package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// )

// const (
// 	position  = "0"
// 	immediate = "1"
// 	relative  = "2"
// )

// type VM struct {
// 	io      string
// 	camfeed func() []int
// 	display func(int)
// }

// type Feeder interface {
// 	Feed() int
// }

// type Broadcaster interface {
// 	Broadcast(int)
// }

// func (vm *VM) Feed() int {
// 	return vm.camfeed()
// }

// func (vm *VM) Broadcast(output int) {
// 	vm.display(output)
// }

// func (vm *VM) Setup(io string, feed func() []int, output func(int)) {
// 	vm.io = io
// 	vm.camfeed = feed
// 	vm.display = output
// }

// var args = map[int]int{1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2}

// func parse(input string) []int {
// 	data, _ := ioutil.ReadFile(input)
// 	scodes := strings.Split(strings.TrimSpace(string(data)), ",")
// 	intcodes := make([]int, len(scodes))
// 	for i, s := range scodes {
// 		intcodes[i], _ = strconv.Atoi(s)
// 	}
// 	return intcodes
// }

// func exec(intcodes []int, input Feeder, notifier Broadcaster, cursor int) []int {
// 	var opcode, inputpos, relativebase int
// 	var modes map[int]string

// 	codes := make([]int, len(intcodes)*10)
// 	copy(codes, intcodes)

// 	for {

// 		sopcode := fmt.Sprintf("%05d", codes[cursor])
// 		opcode, _ = strconv.Atoi(sopcode[3:5])

// 		if opcode == 99 {
// 			break
// 		}

// 		modes = map[int]string{
// 			1: sopcode[2:3],
// 			2: sopcode[1:2],
// 			3: sopcode[0:1],
// 		}

// 		read := func(pos int) int {
// 			switch modes[pos] {
// 			case position:
// 				return codes[codes[cursor+pos]]
// 			case immediate:
// 				return codes[cursor+pos]
// 			case relative:
// 				return codes[relativebase+codes[cursor+pos]]
// 			}
// 			panic("@Read: Invalid operation")
// 		}

// 		write := func(pos int, val int) {
// 			switch modes[pos] {
// 			case position:
// 				codes[codes[cursor+pos]] = val
// 			case relative:
// 				codes[relativebase+codes[cursor+pos]] = val
// 			}
// 		}

// 		switch opcode {
// 		case 1:
// 			write(3, read(1)+read(2))
// 			cursor += args[opcode]
// 		case 2:
// 			write(3, read(1)*read(2))
// 			cursor += args[opcode]
// 		case 3:
// 			write(1, input.Feed())
// 			cursor += args[opcode]
// 		case 4:
// 			notifier.Broadcast(read(1))
// 			cursor += args[opcode]
// 		case 5:
// 			if read(1) != 0 {
// 				cursor = read(2)
// 				continue
// 			}
// 			cursor += args[opcode]
// 		case 6:
// 			if read(1) == 0 {
// 				cursor = read(2)
// 				continue
// 			}
// 			cursor += args[opcode]
// 		case 7:
// 			if read(1) < read(2) {
// 				write(3, 1)
// 			} else {
// 				write(3, 0)
// 			}
// 			cursor += args[opcode]
// 		case 8:
// 			if read(1) == read(2) {
// 				write(3, 1)
// 			} else {
// 				write(3, 0)
// 			}
// 			cursor += args[opcode]
// 		case 9:
// 			relativebase += read(1)
// 			cursor += args[opcode]
// 		}

// 	}
// 	return codes
// }

// func partI(filename string) {
// 	setup := new(VM)
// 	setup.Setup(filename, func() []int {
// 		return [1
// 	}, func(output int) {
// 		fmt.Println(output)
// 	})
// 	Run(setup)()
// }

// func Run(setup *VM) func() {
// 	scodes := parse(setup.io)
// 	return func() {
// 		_ = exec(scodes, setup, setup, 0)
// 	}
// }

// func main() {
// 	partI("day-9-input.txt")
// }
