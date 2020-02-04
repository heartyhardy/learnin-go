package main

import (
	"fmt"
	"time"
)

const matrixSize = 10

type grid [][]*cell

type cell struct {
	x, y, v int
	visited bool
	stacked bool
}

var directions = map[int]cell{
	0: cell{x: 0, y: -1},
	1: cell{x: -1, y: 0},
	2: cell{x: 0, y: 1},
	3: cell{x: 1, y: 0},
}

func pop(stack []*cell) (c *cell, stk []*cell) {
	if len(stack) <= 0 {
		return nil, stack
	}
	c = stack[len(stack)-1]
	stack[len(stack)-1] = nil
	stk = stack[:len(stack)-1]
	return
}

func peek(stack []*cell) (c *cell) {
	if len(stack) <= 0 {
		return nil
	}
	c = stack[len(stack)-1]
	return
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
	stack := make([]*cell, 1)
	(*g)[0][0].stacked = true
	stack[0] = (*g)[0][0]

	// If stack has items
	// pop the top item
	// pick random neighbor
	// mark cell as visited
	// push the neighbor to stack
	// repeat until exhausted

	for len(stack) > 0 {
		next := peek(stack)
		if c, ok := pickNeighbor(g, next); ok {
			c.stacked = true
			stack = append(stack, c)
			(*g)[next.y][next.x].visited = true
			(*g)[next.y][next.x].v = 9
			c.v = 8
			printGrid(g)
			c.v = 9
			time.Sleep(250 * time.Millisecond)
		} else {
			//backtrack
			break
		}
	}

}

func pickNeighbor(g *grid, c *cell) (picked *cell, ok bool) {
	for _, v := range directions {
		dx := v.x + c.x
		dy := v.y + c.y
		if dx >= 0 && dx < matrixSize &&
			dy >= 0 && dy < matrixSize && !c.visited {
			if !(*g)[dy][dx].visited {
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
			case 9:
				fmt.Printf(" \u001b[38;5;208m%v\u001b[0m ", "◼")
			case 8:
				fmt.Printf(" \u001b[38;5;27m%v\u001b[0m ", "◼")
			default:
				fmt.Printf(" %v ", "◽")
			}
		}
		fmt.Println()
	}
}

func main() {
	g := createGrid()
	runDFS(&g)
}
