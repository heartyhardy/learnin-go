package main

import (
	"fmt"
	"time"
)

const matrixSize = 10

const (
	unexplored = iota
	exploring
	explored
	backtracking
)

type grid [][]*cell

type cell struct {
	x, y, v int
	visited bool
}

type xy struct {
	x int
	y int
}

type stack []cell

func (s stack) peek() cell {
	return s[len(s)-1]
}

func (s *stack) pop() (cell, stack) {
	c := (*s)[len(*s)-1]
	(*s)[len(*s)-1].x = 0
	(*s)[len(*s)-1].y = 0
	(*s)[len(*s)-1].v = 0
	(*s)[len(*s)-1].visited = false
	tmp := make([]cell, len(*s))
	copy(tmp, (*s)[:len(*s)-1])
	return c, (*s)[:len(*s)-1]
}

var directions = map[int]cell{
	0: cell{x: 0, y: -1},
	1: cell{x: -1, y: 0},
	2: cell{x: 0, y: 1},
	3: cell{x: 1, y: 0},
}

func createGrid() grid {
	g := make(grid, matrixSize)
	for r := 0; r < matrixSize; r++ {
		g[r] = make([]*cell, matrixSize)
		for c := 0; c < matrixSize; c++ {
			newCell := new(cell)
			newCell.x, newCell.y, newCell.v = c, r, 0
			g[r][c] = newCell
		}
	}
	return g
}

func runDFS(g *grid) {
	cellmap := make(map[xy]bool)
	cellstack := make(stack, 1)
	start := *(*g)[0][0]
	cellmap[xy{x: start.x, y: start.y}] = true
	cellstack = append(cellstack, start)
	// If stack has items
	// pop the top item
	// pick random neighbor
	// mark cell as visited
	// push the neighbor to stack
	// repeat until exhausted
	for len(cellstack) > 0 {
		next := cellstack.peek()

		if v, ok := getAdjacent(g, cellmap, next); ok {
			cellstack = append(cellstack, *v)
			(*g)[next.y][next.x].v = explored
			(*g)[v.y][v.x].v = exploring
			printGrid(g)
			(*g)[v.y][v.x].v = explored
			time.Sleep(250 * time.Millisecond)
		} else if !ok {
			var popped cell
			popped, cellstack = cellstack.pop()
			(*g)[popped.y][popped.x].v = backtracking
			printGrid(g)
			time.Sleep(100 * time.Millisecond)
		}
	}

}

func getAdjacent(g *grid, mstack map[xy]bool, c cell) (*cell, bool) {
	for _, d := range directions {
		dx := c.x + d.x
		dy := c.y + d.y

		if dx >= 0 && dx < matrixSize &&
			dy >= 0 && dy < matrixSize {
			t := xy{x: dx, y: dy}
			if _, ok := mstack[t]; !ok {
				mstack[t] = true
				return (*g)[dy][dx], true
			}
		}
	}
	return nil, false
}

func printGrid(g *grid) {
	fmt.Printf("\033[2J")
	for _, rows := range *g {
		for _, col := range rows {
			switch col.v {
			case explored:
				fmt.Printf(" \u001b[38;5;208m%v\u001b[0m ", "◼")
			case exploring:
				fmt.Printf(" \u001b[38;5;27m%v\u001b[0m ", "◼")
			case backtracking:
				fmt.Printf(" \u001b[38;5;112m%v\u001b[0m ", "◼")
			default:
				fmt.Printf(" %v ", "◽")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	g := createGrid()
	runDFS(&g)

}
