package main

import (
	"heartyhardy/ivm"
)

const (
	black = iota
	white
)

const (
	up = iota
	left
	right
	down
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

type Robot struct {
	panels    map[Panel]int
	state     int
	direction Vector
	x         int
	y         int
}

func (rb *Robot) Paint(color int) {

}

func (rb *Robot) Move(direction int) int {

}

func (rb *Robot) Execute(setup *ivm.VM) func() {

}

//Works as the input feed for VM
func (rb *Robot) Feed() int {
	standing := Panel{x: rb.x, y: rb.y}
	if _, ok := rb.panels[standing]; ok {
		return rb.panels[standing]
	}
	return black
}

func (rb *Robot) Turn(direction int) Panel {
	if v, ok := directions[rb.direction]; ok {
		switch direction {
		case 0:

		}
	}
}
