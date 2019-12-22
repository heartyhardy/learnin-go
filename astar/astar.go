package main

import (
	"fmt"
	"math"
)

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

//GetDistance calculates the manhattan distance to the destination
func (a *AStar) GetDistance(n Node) int {
	return int(math.Abs(float64(n.x-a.destination.x)) +
		math.Abs(float64(n.y-a.destination.y)))
}

//calcH calculates the heuristics for the given grid
func (a *AStar) calcH() {
	mhd := 0
	for i := 0; i < len(a.grid); i++ {
		for j := 0; j < len(a.grid); j++ {
			mhd = a.GetDistance(a.grid[i][j])
			a.grid[i][j].h = mhd
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
			if i == a.origin.x && j == a.origin.y {
				fmt.Print(fmt.Sprintf(" %2s ", "o"))
			} else if i == a.destination.x && j == a.destination.y {
				fmt.Print(fmt.Sprintf(" %2s ", "x"))
			} else if j == 9 {
				fmt.Print("\n")
			} else {
				//fmt.Print(" . ")
				fmt.Print(fmt.Sprintf(" %2d ", a.grid[i][j].h))
			}
		}
	}
}

func main() {
	a := AStar{}
	a.Populate(10)
	a.SetOrigin(1, 1)
	a.SetDestination(6, 8)
	a.calcH()
	a.PrintGrid()
}
