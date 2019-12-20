package main

import (
	"fmt"
	"heartyhardy/ivm"
)

//Direction flags
const (
	_ = iota
	north
	south
	west
	east
)

//Executor carries out setting up and executing ivm
type Executor interface {
	Execute(ivm.VM) func()
}

//Responser receives output from ivm
type Responser interface {
	Responser(int)
}

//Requester sends input to ivm
type Requester interface {
	Request() []int
}

//Droid struct
type Droid struct {
}

//Command struct stores info on a Droid command
// base : Base command ex: N / S / W / E
// inverse direction of a given command ex: N -> S, E -> W
// length : number of explored routes
// routes by index [true, false, true, false] = 0 index is open, 1st index is a wall so on (index + 1 is a Command)
type Command struct {
	base        int
	inverse     int
	length      int
	routes      []bool
	backtracked bool
}

//Request sends a message as input to the ivm
func (d *Droid) Request() []int {
	return []int{2, 2, 4, 4, 4, 4, 4, 4, 4, 4, 2, 2}
}

//Response collects the output messages from ivm
func (d *Droid) Response(message int) {
	fmt.Print(message, ", ")
}

//Execute runs the ivm with specified setup instructions
func (d *Droid) Execute(setup *ivm.VM) func() {
	return ivm.Run(setup)
}

//PartI ...
func PartI() func() {
	puzzleInput := "ivm-droid-input.txt"
	d := new(Droid)
	vm := new(ivm.VM)
	vm.Setup(puzzleInput, d.Request, d.Response)
	exec := d.Execute(vm)
	return exec
}

func main() {
	PartI()()
}
