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
