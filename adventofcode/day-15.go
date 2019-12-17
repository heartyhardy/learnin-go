package main

// import (
// 	"fmt"
// 	"heartyhardy/ivm"
// )

// //Direction flags
// const (
// 	_ = iota
// 	north
// 	south
// 	west
// 	east
// )

// //Bitwise flags used in proximity scan
// const (
// 	N = 1 << iota
// 	S
// 	W
// 	E
// )

// //Droid states
// const (
// 	idle = iota
// 	listening
// 	moving
// 	backtracking
// )

// //Executor carries out setting up and executing ivm
// type Executor interface {
// 	Execute(ivm.VM) func()
// }

// //Responser receives output from ivm
// type Responser interface {
// 	Responser(int)
// }

// //Requester sends input to ivm
// type Requester interface {
// 	Request() []int
// }

// //Droid ...
// // Sitrep stores bitwise flag for proximity sacn
// // lookup stores the current direction
// // base is the current base direction of the Droids pos ex: N -> E -> S -> N is the current Base
// // routes stores all the possible routes found. Each key denotes a subroute if any
// type Droid struct {
// 	Lookup
// 	routes   map[int][]int
// 	requests []int
// 	base     int
// 	sitrep   int
// 	status   int
// }

// //Lookup ...
// type Lookup struct {
// 	lookups     []int
// 	lookup      int
// 	lookupindex int
// }

// //Gets the flag for given direction
// func getFlag(flag int) int {
// 	switch flag {
// 	case 1:
// 		return N
// 	case 2:
// 		return S
// 	case 3:
// 		return W
// 	case 4:
// 		return E
// 	default:
// 		fmt.Println("INVALID:", flag)
// 		panic("Invalid flag")
// 	}
// }

// //SetScan saves a scan direction status in each bit
// func (d *Droid) SetScan(flag int) {
// 	d.sitrep |= flag
// }

// //GetScan retrieves direction status by flag
// func (d *Droid) GetScan(flag int) int {
// 	return d.sitrep & flag
// }

// //SitRep Gets the set flags for each direction denoting possible routes
// func (d *Droid) SitRep() ([]int, bool) {
// 	direction := north
// 	nopath := false
// 	sr := make([]int, 0)
// 	for i := 1; i < 9; i = i * 2 {
// 		nopath = nopath || d.sitrep&i != 0
// 		if d.sitrep&i != 0 {
// 			sr = append(sr, direction)
// 		}
// 		direction++
// 	}
// 	return sr, nopath
// }

// //AddRequest adds a direction to scan when Request method triggers
// func (d *Droid) AddRequest(req int) {
// 	if req > 0 && req < 5 {
// 		d.requests = append(d.requests, req)
// 	}
// }

// //GetLookup gets a valid direction to move
// func (d *Droid) GetLookup(lookup int) []int {
// 	lookups := map[int][]int{
// 		0: []int{1, 2, 3, 4},
// 		1: []int{1, 3, 4},
// 		2: []int{2, 3, 4},
// 		3: []int{1, 2, 3},
// 		4: []int{1, 2, 4},
// 	}
// 	return lookups[lookup]
// }

// //IsDeadEnd checks for a possible deadend
// func (d *Droid) IsDeadEnd() bool {
// 	if _, ok := d.SitRep(); !ok && d.lookupindex > 2 {
// 		return true
// 	}
// 	return false
// }

// //TurnScanner scans a direction after rotating the scanner
// func (d *Droid) TurnScanner() int {

// 	if d.lookupindex > len(d.GetLookup(d.lookup))-1 {
// 		d.status = backtracking
// 	}
// 	scandir := d.GetLookup(d.lookup)[d.lookupindex]
// 	return scandir
// }

// //ProximityScan checks out adjacent directions for pathables
// func (d *Droid) ProximityScan() {
// 	switch d.status {
// 	case moving:
// 		d.base = d.requests[len(d.requests)-1]
// 		d.lookup = d.base
// 		d.lookupindex = 0
// 		d.lookup = d.TurnScanner()
// 		d.AddRequest(d.lookup)
// 		d.status = idle
// 	case idle:
// 		d.lookupindex++
// 		scandir := d.TurnScanner()
// 		d.AddRequest(scandir)
// 		d.status = listening
// 	}
// }

// //Request Implementing ivm Feed for inputs
// func (d *Droid) Request() []int {
// 	d.ProximityScan()
// 	fmt.Println("SENT: ", d.requests[len(d.requests)-1], " ")
// 	//a := []int{2, 2, 4, 4, 4, 4, 4, 4, 4, 4, 4, 2, 2, 2, 2, 3, 3, 1, 1, 3}
// 	return d.requests
// }

// //Response captures output from ivm
// func (d *Droid) Response(message int) {
// 	fmt.Println("GOT: ", message)
// 	switch message {
// 	case 0:
// 		d.status = idle
// 	case 1:
// 		//fmt.Print(d.lookup, ",")
// 		d.status = moving
// 	case 2:
// 		panic("FOUND IT!")
// 	}
// }

// //Execute runs the ivm with specified setup instructions
// func (d *Droid) Execute(setup *ivm.VM) func() {
// 	return ivm.Run(setup)
// }

// //Boot initializes the Droid
// func (d *Droid) Boot() {
// 	d.lookup = 0
// }

// //PartI ...
// func PartI() func() {
// 	puzzleInput := "ivm-droid-input.txt"
// 	d := new(Droid)
// 	d.Boot()
// 	d.routes = make(map[int][]int)
// 	vm := new(ivm.VM)
// 	vm.Setup(puzzleInput, d.Request, d.Response)
// 	exec := d.Execute(vm)
// 	return exec
// }

// func main() {
// 	PartI()()
// }
