package aoc15

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const matrixSize = 100

type lightState [][]*state

type state struct {
	current int
	next    int
}

type rc struct {
	r int
	c int
}

var directions = map[int]rc{
	0: rc{r: -1, c: 0},
	1: rc{r: -1, c: 1},
	2: rc{r: 0, c: 1},
	3: rc{r: 1, c: 1},
	4: rc{r: 1, c: 0},
	5: rc{r: 1, c: -1},
	6: rc{r: 0, c: -1},
	7: rc{r: -1, c: -1},
}

func readLightData(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func setInitialState(lightdata []string, isGameOfLife bool) lightState {
	lights := make(lightState, matrixSize)

	for r, rows := range lightdata {
		lights[r] = make([]*state, matrixSize)
		for c, col := range rows {
			lstate := new(state)
			switch col {
			case 46:
				lstate.current = 0
			case 35:
				lstate.current = 1
			}
			lights[r][c] = lstate
		}
	}

	if isGameOfLife {
		setRestrictions(&lights)
	}

	return lights
}

func setRestrictions(lights *lightState) {
	(*lights)[0][0].current = 1
	(*lights)[0][matrixSize-1].current = 1
	(*lights)[matrixSize-1][0].current = 1
	(*lights)[matrixSize-1][matrixSize-1].current = 1
}

func letThereBeLight(lights *lightState, steps int, isGameOfLife bool) {
	for t := 0; t < steps; t++ {
		for r, rows := range *lights {
			for c, col := range rows {

				if isGameOfLife {
					if isRestricted(r, c) {
						continue
					}
				}

				on, _ := getAdjacentStatus(lights, r, c)
				switch col.current {
				case 0:
					if on == 3 {
						(*lights)[r][c].next = 1
						continue
					}
					(*lights)[r][c].next = (*lights)[r][c].current
				case 1:
					if on != 2 && on != 3 {
						(*lights)[r][c].next = 0
						continue
					}
					(*lights)[r][c].next = (*lights)[r][c].current
				}
			}
		}
		consoleOut(lights, isGameOfLife)
	}
}

func consoleOut(lights *lightState, isGameOfLife bool) {
	for r, rows := range *lights {
		for c, col := range rows {

			if isGameOfLife {
				if !isRestricted(r, c) {
					(*lights)[r][c].current, (*lights)[r][c].next = (*lights)[r][c].next, (*lights)[r][c].current
				}
			} else {
				(*lights)[r][c].current, (*lights)[r][c].next = (*lights)[r][c].next, (*lights)[r][c].current
			}

			switch col.current {
			case 0:
				fmt.Print("\u001b[30m⚯\u001b[0m")
			case 1:
				fmt.Print("\u001b[36;1m⊷\u001b[0m")
			}
		}
		fmt.Println("")
	}
}

func countOnLights(lights *lightState) int {
	count := 0
	for _, rows := range *lights {
		for _, col := range rows {
			if col.current == 1 {
				count++
			}
		}
	}
	return count
}

func getAdjacentStatus(lights *lightState, r, c int) (int, int) {
	on, off := 0, 0
	for _, dv := range directions {
		dr := r + dv.r
		dc := c + dv.c
		if (dr < matrixSize && dr >= 0) && (dc < matrixSize && dc >= 0) {
			if (*lights)[dr][dc].current == 1 {
				on++
				continue
			}
			off++
		}
	}
	return on, off
}

func isRestricted(r, c int) bool {
	switch {
	case r == 0 && c == 0:
		return true
	case r == 0 && c == matrixSize-1:
		return true
	case r == matrixSize-1 && c == 0:
		return true
	case r == matrixSize-1 && c == matrixSize-1:
		return true
	default:
		return false
	}

}

//Run Solution
func Run() {
	lightData := readLightData("./inputs/day-18.txt")
	lights := setInitialState(lightData, false)
	letThereBeLight(&lights, 100, false)
	beforeGameofLight := countOnLights(&lights)
	fmt.Println("\n\nPART I: \n", beforeGameofLight)
	fmt.Println("Resuming in 5...")
	time.Sleep(5 * time.Second)

	lights = setInitialState(lightData, true)
	letThereBeLight(&lights, 100, true)
	gameOfLife := countOnLights(&lights)
	fmt.Println("\n\nPART II: \n", gameOfLife)

}
