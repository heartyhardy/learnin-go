// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// )

// const (
// 	positionMode  = "0"
// 	immediateMode = "1"
// 	relativeMode  = "2"
// )

// // Current Directions
// const (
// 	U = "U"
// 	D = "D"
// 	L = "L"
// 	R = "R"
// )

// //Status of the HullPainter
// const (
// 	_ = iota
// 	Idle
// 	Painting
// 	Moving
// )

// //Grid used for storing coordinates of the robot
// type Grid struct {
// 	X int
// 	Y int
// }

// //GridProps stores Grid properties
// type GridProps struct {
// 	Color      int
// 	Paintcount int
// }

// var arglens map[int]int = map[int]int{1: 4, 2: 4, 3: 2, 4: 2, 5: 3, 6: 3, 7: 4, 8: 4, 9: 2}

// var directions = map[string]map[string]Grid{
// 	"U": map[string]Grid{
// 		"L": Grid{-1, 0},
// 		"R": Grid{1, 0},
// 	},
// 	"D": map[string]Grid{
// 		"L": Grid{1, 0},
// 		"R": Grid{-1, 0},
// 	},
// 	"L": map[string]Grid{
// 		"L": Grid{0, -1},
// 		"R": Grid{0, 1},
// 	},
// 	"R": map[string]Grid{
// 		"L": Grid{0, 1},
// 		"R": Grid{0, -1},
// 	},
// }

// var face = map[string]map[string]string{
// 	"U": map[string]string{
// 		"L": "L",
// 		"R": "R",
// 	},
// 	"D": map[string]string{
// 		"L": "R",
// 		"R": "L",
// 	},
// 	"L": map[string]string{
// 		"L": "D",
// 		"R": "U",
// 	},
// 	"R": map[string]string{
// 		"L": "U",
// 		"R": "D",
// 	},
// }

// //HullPainter Robot
// type HullPainter struct {
// 	surface      map[Grid]GridProps
// 	location     Grid
// 	facing       string
// 	status       int
// 	defaultColor int
// }

// //Dtoa Converts int to string
// func (hp *HullPainter) Dtoa(facing int) string {
// 	switch facing {
// 	case 0:
// 		return L
// 	case 1:
// 		return R
// 	}
// 	panic("@HullPainter: Dtoa: " + string(facing))
// }

// //Process passes intcode output instructions to the Robot
// func (hp *HullPainter) Process(code int) {
// 	switch hp.status {
// 	case Idle:
// 		hp.Paint(code)
// 	case Painting:
// 		hp.Move(code)
// 	case Moving:
// 	default:
// 		panic("@HullPainter: Process: " + string(hp.status))
// 	}

// }

// //Paint - Paints a given Grid with the color
// func (hp *HullPainter) Paint(color int) {
// 	hp.status = Painting
// 	if _, ok := hp.surface[hp.location]; ok {
// 		surface := hp.surface[hp.location]
// 		newProps := GridProps{
// 			Color:      color,
// 			Paintcount: surface.Paintcount + 1,
// 		}
// 		hp.surface[hp.location] = newProps
// 	}
// }

// //Move sets the robots moving direction and moves it to a new location
// func (hp *HullPainter) Move(code int) {
// 	hp.status = Moving
// 	direction := hp.Dtoa(code)
// 	nextDirection := face[hp.GetFacing()][direction]
// 	offset := directions[hp.GetFacing()][direction]
// 	grid := Grid{
// 		X: hp.location.X + offset.X,
// 		Y: hp.location.Y + offset.Y,
// 	}
// 	hp.location = grid
// 	hp.facing = nextDirection

// 	if _, ok := hp.surface[hp.location]; !ok {
// 		hp.surface[hp.location] = GridProps{Color: 0, Paintcount: 0}
// 	}

// 	hp.status = Idle
// }

// //GetFacing returns the current direction robot is facing
// func (hp *HullPainter) GetFacing() string {
// 	return hp.facing
// }

// //GetLocation returns the current location of the robot
// func (hp *HullPainter) GetLocation() Grid {
// 	return hp.location
// }

// //GetSurface returns the Hull paint surface
// func (hp *HullPainter) GetSurface() map[Grid]GridProps {
// 	return hp.surface
// }

// //GetColor returns the color of hhe current Grid
// func (hp *HullPainter) GetColor() int {
// 	if _, ok := hp.surface[hp.location]; ok {
// 		props := hp.surface[hp.location]
// 		return props.Color
// 	}
// 	return hp.defaultColor
// }

// func exec(src []int, args []int, cursor int) (flush []int, exitcode int) {
// 	exitcode = 1
// 	const bufferSize = 3000
// 	relativeBase := 0
// 	codes := make([]int, bufferSize)
// 	flush = make([]int, 0)
// 	copy(codes, src)

// 	hullPainter := new(HullPainter)
// 	hullPainter.surface = make(map[Grid]GridProps)
// 	hullPainter.location = Grid{X: 0, Y: 0}
// 	hullPainter.surface[hullPainter.location] = GridProps{0, 0}
// 	hullPainter.facing = U
// 	hullPainter.status = Idle

// 	for {
// 		fmtcode := fmt.Sprintf("%05d", codes[cursor])
// 		opcode, _ := strconv.Atoi(fmtcode[3:5])

// 		if opcode == 99 {
// 			break
// 		}

// 		modes := map[int]string{
// 			1: fmtcode[2:3],
// 			2: fmtcode[1:2],
// 			3: fmtcode[0:1],
// 		}

// 		peek := func(pos int) int {
// 			switch modes[pos] {
// 			case positionMode:
// 				return codes[codes[cursor+pos]]
// 			case immediateMode:
// 				return codes[cursor+pos]
// 			case relativeMode:
// 				return codes[relativeBase+codes[cursor+pos]]
// 			}
// 			panic("(@peek) Invalid mode: " + modes[pos])
// 		}

// 		save := func(pos int, val int) {
// 			switch modes[pos] {
// 			case positionMode:
// 				codes[codes[cursor+pos]] = val
// 				return
// 			case relativeMode:
// 				codes[relativeBase+codes[cursor+pos]] = val
// 				return
// 			}
// 			panic("(@save) invalid mode: " + modes[pos])
// 		}

// 		switch opcode {
// 		case 1:
// 			save(3, peek(1)+peek(2))
// 			cursor += arglens[opcode]
// 		case 2:
// 			save(3, peek(1)*peek(2))
// 			cursor += arglens[opcode]
// 		case 3:
// 			save(1, hullPainter.GetColor())
// 			cursor += arglens[opcode]
// 		case 4:
// 			hullPainter.Process(peek(1))
// 			flush = append(flush, peek(1))
// 			cursor += arglens[opcode]
// 		case 5:
// 			if peek(1) != 0 {
// 				cursor = peek(2)
// 			} else {
// 				cursor += arglens[opcode]
// 			}
// 		case 6:
// 			if peek(1) == 0 {
// 				cursor = peek(2)
// 			} else {
// 				cursor += arglens[2]
// 			}
// 		case 7:
// 			if peek(1) < peek(2) {
// 				save(3, 1)
// 			} else {
// 				save(3, 0)
// 			}
// 			cursor += arglens[opcode]
// 		case 8:
// 			//fmt.Println("PEEK 8", peek(1), peek(2), codes[cursor+3], len(src))
// 			if peek(1) == peek(2) {
// 				save(3, 1)
// 			} else {
// 				save(3, 0)
// 			}
// 			cursor += arglens[opcode]
// 		case 9:
// 			relativeBase += peek(1)
// 			cursor += arglens[opcode]
// 		default:
// 			panic("Invalid opcode: " + string(opcode))
// 		}
// 	}
// 	return flush, 0
// }

// func parse(input string) (intCodes []int) {
// 	scodes := strings.Split(strings.ReplaceAll(input, "\r\n", ""), ",")
// 	intCodes = make([]int, len(scodes))
// 	for i, str := range scodes {
// 		intCodes[i], _ = strconv.Atoi(str)
// 	}
// 	return
// }

// func boostkey(input string) string {
// 	intCodes := parse(string(input))
// 	output, _ := exec(intCodes, []int{1}, 0)
// 	if len(output) != 1 {
// 		fmt.Printf("OUTPUT: %v/n", output)
// 	}
// 	return fmt.Sprintf("%s%s", "Key: ", strconv.Itoa(output[len(output)-1]))
// }

// func sos(input string) string {
// 	intCodes := parse(string(input))
// 	output, _ := exec(intCodes, []int{2}, 0)
// 	if len(output) != 1 {
// 		fmt.Printf("OUTPUT: %v/n", output)
// 	}
// 	return fmt.Sprintf("%s%s", "SOS: ", strconv.Itoa(output[len(output)-1]))
// }

// func runrobo(input string) {
// 	intCodes := parse(input)
// 	_, _ = exec(intCodes, []int{0}, 0)

// }

// func main() {
// 	input, _ := ioutil.ReadFile("day-11-input.txt")
// 	runrobo(string(input))

// }
