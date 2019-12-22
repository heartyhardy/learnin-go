package main

import "fmt"

//Node is used to store grid data for A*
type Node struct {
	g      int
	h      int
	f      int
	x      int
	y      int
	parent *Node
}

//AStar ...
type AStar struct {
	grid        [][]Node
	origin      *Node
	destination *Node
}

//Populate the 2D grid
func (a *AStar) Populate(gsize int) {
	a.grid = make([][]Node, gsize)
	for i := 0; i < gsize; i++ {
		a.grid[i] = make([]Node, gsize)
		for j := 0; j < gsize; j++ {
			a.grid[i][j] = Node{x: i, y: j, parent: nil}
		}
	}
}

//SetOrigin sets the starting point of the grid
func (a *AStar) SetOrigin(x, y int) {
	a.origin = &a.grid[y][x]
}

//SetDestination sets the end point of the grid
func (a *AStar) SetDestination(x, y int) {
	a.destination = &a.grid[y][x]
}

//PrintGrid ...
func (a *AStar) PrintGrid() {
	for i := 0; i < len(a.grid); i++ {
		for j := 0; j < len(a.grid); j++ {
			fmt.Print(" . ")
			if j == 9 {
				fmt.Print("\n")
			}
		}
	}
}

func main() {
	a := AStar{}
	a.Populate(10)
	a.PrintGrid()
}
