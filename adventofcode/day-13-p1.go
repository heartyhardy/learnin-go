package main

import (
	"fmt"
	"heartyhardy/ivm"
)

const (
	empty = iota
	wall
	block
	paddle
	ball
)

const (
	idle = iota
	xReady
	yReady
	idReady
)

type TileBuilder interface {
	BuildTile(int, int, int) Tile
}

type Scorer interface {
	Score(int)
}

type Receiver interface {
	Receive(int)
}

type Executor interface {
	Execute(ivm.VM) func()
}

type TileProps struct {
	id int
}

type Tile struct {
	x int
	y int
}

type Buffer struct {
	x  int
	y  int
	id int
}

type Game struct {
	messages      []int
	tiles         map[Tile]TileProps
	receiveBuffer *Buffer
	state         int
	score         int
	ballPos       int
	paddlePos     int
}

func (bf *Buffer) Get() *Buffer {
	return bf
}

func (bf *Buffer) Put(mode int, data int) {
	switch mode {
	case idle:
		bf.x = data
	case xReady:
		bf.y = data
	case yReady:
		bf.id = data
	}
}

func (bf *Buffer) Flush(g *Game) {

	if bf.x == -1 && bf.y == 0 && (bf.id < 0 || bf.id > 4) {
		g.Score(bf.id)
		g.state = idle
	} else {
		switch bf.Get().id {
		case empty:
			fallthrough
		case wall:
			fallthrough
		case block:
			g.BuildTile(bf.Get().x, bf.Get().y, bf.Get().id)
		case 3:
			g.paddlePos = bf.x
			g.state = idle
		case 4:
			g.ballPos = bf.x
			g.state = idle
		}
	}
}

func (g *Game) Feed() []int {

	joy := 0
	switch {
	case g.ballPos < g.paddlePos:
		joy = -1
	case g.ballPos > g.paddlePos:
		joy = 1
	default:
		joy = 0
	}

	g.messages = append(g.messages, joy)
	return g.messages
}

func (g *Game) Receive(message int) {
	switch g.state {
	case idle:
		g.receiveBuffer.Put(idle, message)
		g.state = xReady
	case xReady:
		g.receiveBuffer.Put(xReady, message)
		g.state = yReady
	case yReady:
		g.receiveBuffer.Put(yReady, message)
		g.state = idReady
		g.receiveBuffer.Flush(g)
	}
}

func (g *Game) BuildTile(x int, y int, id int) *Tile {
	t := &Tile{x: x, y: y}
	tp := &TileProps{id: id}
	g.tiles[*t] = *tp
	g.state = idle

	return t
}

func (g *Game) Score(sco int) {
	g.score = sco
}

func (g *Game) Execute(setup *ivm.VM) func() {
	return ivm.Run(setup)
}

func initialize() *Game {
	g := new(Game)
	g.receiveBuffer = new(Buffer)
	g.messages = make([]int, 0)
	g.tiles = make(map[Tile]TileProps)
	g.state = idle
	return g
}

func PartI() {
	puzzleInput := "day-13-input.txt"
	g := initialize()
	vm := new(ivm.VM)
	vm.Setup(puzzleInput, g.Feed, g.Receive)
	exec := g.Execute(vm)
	exec()

	blocks, walls, empties := 0, 0, 0
	for _, tile := range g.tiles {
		switch tile.id {
		case empty:
			empties++
		case wall:
			walls++
		case block:
			blocks++
		}
	}

	fmt.Println("\n\u001b[34m------------------- TILE STATS -----------------------------\u001b[0m")
	fmt.Printf("\t\t \u001b[32mBLOCK TILES: %v\u001b[0m                   \n", blocks)
	fmt.Printf("\t\t \u001b[32mWALL TILES:   %v \u001b[0m                   \n", walls)
	fmt.Printf("\t\t \u001b[32mEMPTY TILES: %v \u001b[0m                   \n", empties)
	fmt.Println("\u001b[34m------------------------------------------------------------\u001b[0m")
	fmt.Println()

}

func PartII() {
	puzzleInput := "day-13-play.txt"
	g := initialize()
	vm := new(ivm.VM)
	vm.Setup(puzzleInput, g.Feed, g.Receive)
	exec := g.Execute(vm)
	exec()

	fmt.Println("\n\u001b[31;1m------------------- GAME OVER! -----------------------------\u001b[0m\n")
	fmt.Printf("\t\t \u001b[32mSCORE: %v\u001b[0m                   \n\n", g.score)
	fmt.Println("\u001b[33;1m------------------------------------------------------------\u001b[0m")
	fmt.Println()
}

func main() {
	PartI()
	PartII()
}
