package main

import (
	"fmt"
	"heartyhardy/ivm"
)

var count int

const (
	black = iota
	white
)

const (
	paintReady = iota
	moveReady
)

type Vector struct {
	x int
	y int
}

func (v *Vector) Flip() *Vector {
	t := v.x
	v.x = v.y
	v.y = t
	return v
}

func (v *Vector) Advance(w *Vector) *Vector {
	v.x = v.x + w.x
	v.y = v.y + w.y
	return v
}

//Flip key and multiply by value
var directions map[Vector]Vector = map[Vector]Vector{
	Vector{0, 1}:  Vector{-1, 1},
	Vector{-1, 0}: Vector{1, -1},
	Vector{0, -1}: Vector{-1, 1},
	Vector{1, 0}:  Vector{1, -1},
}

type Panel struct {
	x int
	y int
}

type Painter interface {
	Paint(int)
}

type Mover interface {
	Move(int) int
}

type Executor interface {
	Execute(ivm.VM) func()
}

type Receiver interface {
	Receive(int)
}

type Robot struct {
	messages  []int
	panels    map[Panel]int
	state     int
	direction *Vector
	location  *Vector
}

func (rb *Robot) Paint(color int) {
	panel := Panel{x: rb.location.x, y: rb.location.y}
	if _, ok := rb.panels[panel]; ok {
		rb.panels[panel] = color
		rb.state = moveReady
	}
}

func (rb *Robot) Move(direction int) {
	if d, ok := rb.Turn(direction); ok {
		pos := rb.location.Advance(d)
		panel := Panel{x: pos.x, y: pos.y}

		if _, exists := rb.panels[panel]; !exists {
			rb.panels[panel] = black
		}
		rb.state = paintReady
	}
}

func (rb *Robot) Execute(setup *ivm.VM) func() {
	return ivm.Run(setup)
}

//Works as the input feed for VM
func (rb *Robot) Feed() []int {
	color := black
	standing := Panel{x: rb.location.x, y: rb.location.y}
	if _, ok := rb.panels[standing]; ok {
		color = rb.panels[standing]
	}
	rb.messages = append(rb.messages, color)
	return rb.messages
}

func (rb *Robot) Receive(message int) {

	switch rb.state {
	case paintReady:
		rb.Paint(message)
	case moveReady:
		rb.Move(message)
	}
}

func (rb *Robot) Turn(direction int) (*Vector, bool) {
	if v, ok := directions[*rb.direction]; ok {
		switch direction {
		case 0:
			rb.direction.Flip()
			rb.direction.x = rb.direction.x * v.x
			rb.direction.y = rb.direction.y * v.x
		case 1:
			rb.direction.Flip()
			rb.direction.x = rb.direction.x * v.y
			rb.direction.y = rb.direction.y * v.y
		}
		return rb.direction, ok
	}
	return rb.direction, false
}

func initialize() *Robot {
	rb := new(Robot)
	rb.direction = new(Vector)
	rb.location = new(Vector)
	panel := new(Panel)
	rb.panels = make(map[Panel]int)
	rb.messages = make([]int, 0)
	rb.location.x = 0
	rb.location.y = 0
	rb.direction.x = 0
	rb.direction.y = 1
	panel.x = rb.location.x
	panel.y = rb.location.y
	rb.panels[*panel] = 0
	return rb
}

func PartI() {
	puzzleInput := "ivm-robot-input.txt"
	rb := initialize()
	vm := new(ivm.VM)
	vm.Setup(puzzleInput, rb.Feed, rb.Receive)
	exec := rb.Execute(vm)
	exec()
	fmt.Println(len(rb.panels))
}

func main() {
	PartI()
}
