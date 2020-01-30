package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	blocked = iota
	pathable
	route
)

type tile struct {
	x        int
	y        int
	v        uint8
	visited  bool
	enqueued bool
	parent   *tile
}

type tileMap [][]*tile

var directions = map[int]tile{
	0: tile{x: 0, y: -1},
	1: tile{x: -1, y: 0},
	2: tile{x: 0, y: 1},
	3: tile{x: 1, y: 0},
}

func loadGridFromFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error occured while loading grid data!")
	}
	content := strings.Split(string(data), "\n")
	return content
}

func buildGrid(content []string) tileMap {
	grid := make(tileMap, len(content))
	for rn, fields := range content {
		rows := strings.Fields(fields)
		grid[rn] = make([]*tile, len(rows))
		for cn, col := range rows {
			uv, err := strconv.ParseUint(col, 10, 8)
			if err != nil {
				panic("Invalid data found in grid!")
			}
			newTile := new(tile)
			newTile.x, newTile.y, newTile.v, newTile.parent = cn, rn, uint8(uv), nil
			grid[rn][cn] = newTile
		}
	}
	return grid
}

func findShortestPath(grid *tileMap, start *tile, dest *tile) {
	//Make a queue
	queue := make([]*tile, 1)
	//Enqueue origin
	queue[0] = start
	//While queue is not empty
	//queue surrounding cells if they are not visited
	// NOTE: Need to reset the grid after running BFS which isnt implemented yet
	found := false

	for !found && len(queue) > 0 {
		tl := queue[0]
		queue = queue[1:]
		found = getAdjacent(grid, &queue, tl, dest)
	}
	backtrack(grid, dest)
}

func getAdjacent(grid *tileMap, queue *[]*tile, tl *tile, dest *tile) (found bool) {
	for _, v := range directions {
		x := tl.x + v.x
		y := tl.y + v.y
		if x >= 0 && x < len((*grid)[0]) &&
			y >= 0 && y < len(*grid) &&
			(!(*grid)[y][x].visited && !(*grid)[y][x].enqueued) {

			switch (*grid)[y][x].v {
			case 1:
				(*grid)[y][x].parent = tl
				(*grid)[y][x].enqueued = true
				*queue = append(*queue, (*grid)[y][x])

				if (*grid)[y][x].x == dest.x && (*grid)[y][x].y == dest.y {
					found = true
					tl.visited = true
					return
				}
			}
		}
	}
	tl.visited = true
	return false
}

func backtrack(grid *tileMap, dest *tile) {
	next := dest
	fmt.Println((*grid)[0][0].parent == nil)
	for next != nil {
		(*grid)[next.y][next.x].v = route
		printGrid(*grid)
		next = next.parent
		time.Sleep(1000 * time.Millisecond)
	}
}

func printGrid(grid tileMap) {
	fmt.Printf("\033[2J")
	for _, rows := range grid {
		for _, col := range rows {
			switch col.v {
			case 0:
				fmt.Printf("\u001b[41m %v \u001b[0m", col.v)
			case 1:
				fmt.Printf("\u001b[44m %v \u001b[0m", col.v)
			case 2:
				fmt.Printf("\u001b[42m %v \u001b[0m", col.v)
			}
		}
		fmt.Println()
	}
}

func main() {
	content := loadGridFromFile("grid.txt")
	grid := buildGrid(content)
	gridBFS := make(tileMap, len(grid))
	copy(gridBFS, grid)
	printGrid(grid)
	time.Sleep(2 * time.Second)

	start := gridBFS[0][0]
	dest := gridBFS[3][7]
	findShortestPath(&gridBFS, start, dest)
	fmt.Println()
}
