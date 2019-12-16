package main

//Direction flags
const (
	_ = iota
	nortth
	south
	west
	east
)

//Bitwise flags used in proximity scan
const (
	N = 1 << iota
	S
	W
	E
)

//Droid ...
// Sitrep stores bitwise flag for proximity sacn
// lookup stores the current direction
// base is the current base direction of the Droids pos ex: N -> E -> S -> N is the current Base
// routes stores all the possible routes found. Each key denotes a subroute if any
type Droid struct {
	routes map[int][]int
	base   int
	lookup int
	sitrep int
}

//SetScan saves a scan direction status in each bit
func (d *Droid) SetScan(flag int) {
	d.sitrep |= flag
}

//GetScan retrieves direction status by flag
func (d *Droid) GetScan(flag int) int {
	return d.sitrep & flag
}

//SitRep Gets the set flags for each direction denoting possible routes
func (d *Droid) SitRep() []int {
	direction := nortth
	sr := make([]int, 0)
	for i := 1; i < 9; i = i * 2 {
		if d.sitrep&i != 0 {
			sr = append(sr, direction)
		}
		direction++
	}
	return sr
}

func (d *Droid) ProximityScan() bool {

}

func main() {

}
