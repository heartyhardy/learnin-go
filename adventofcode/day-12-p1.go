package main

// import (
// 	"fmt"
// 	"math"
// )

// type Influencer interface {
// 	Influence(Moon)
// }

// type Velocitizer interface {
// 	Volitize()
// }

// type EnergyCalculator interface {
// 	PotentialEnergy()
// 	KineticEnergy()
// }

// type Comparator interface {
// 	IsPositionEqual(Moon) bool
// 	IsVelocityEqual(Moon) bool
// }

// type Moon struct {
// 	X  int
// 	Y  int
// 	Z  int
// 	VX int
// 	VY int
// 	VZ int
// }

// func (m *Moon) Influence(n *Moon) {
// 	if m.X > n.X {
// 		m.VX += -1
// 	} else if m.X < n.X {
// 		m.VX += 1
// 	}
// 	if m.Y > n.Y {
// 		m.VY += -1
// 	} else if m.Y < n.Y {
// 		m.VY += 1
// 	}
// 	if m.Z > n.Z {
// 		m.VZ += -1
// 	} else if m.Z < n.Z {
// 		m.VZ += 1
// 	}
// }

// func (m *Moon) Velocitize() {
// 	m.X += m.VX
// 	m.Y += m.VY
// 	m.Z += m.VZ
// }

// func (m *Moon) IsPositionEqual(n *Moon) bool {
// 	if m.X == n.X && m.Y == n.Y && m.Z == n.Z {
// 		return true
// 	}
// 	return false
// }

// func (m *Moon) IsVelocityEqual(n *Moon) bool {
// 	if m.VX == n.VX && m.VY == n.VY && m.VZ == n.VZ {
// 		return true
// 	}
// 	return false
// }

// func (m *Moon) PotentialEnergy() int {
// 	return int(math.Abs(float64(m.X)) + math.Abs(float64(m.Y)) + math.Abs(float64(m.Z)))
// }

// func (m *Moon) KineticEnergy() int {
// 	return int(math.Abs(float64(m.VX)) + math.Abs(float64(m.VY)) + math.Abs(float64(m.VZ)))
// }

// func (m *Moon) ApplyGravity(moons ...*Moon) {
// 	for _, moon := range moons {
// 		m.Influence(moon)
// 	}
// }

// func (m *Moon) ApplyVelocity(moons ...*Moon) {
// 	for _, moon := range moons {
// 		moon.Velocitize()
// 	}
// }

// func (m *Moon) TotalEnergy(moons ...*Moon) int {
// 	totalEnergy := 0
// 	for _, moon := range moons {
// 		totalEnergy += moon.PotentialEnergy() * moon.KineticEnergy()
// 	}
// 	return totalEnergy
// }

// func PartI() {

// 	io := &Moon{X: 3, Y: -6, Z: 6, VX: 0, VY: 0, VZ: 0}
// 	europa := &Moon{X: 10, Y: 7, Z: -9, VX: 0, VY: 0, VZ: 0}
// 	ganymede := &Moon{X: -3, Y: -7, Z: 9, VX: 0, VY: 0, VZ: 0}
// 	callisto := &Moon{X: -8, Y: 0, Z: 4, VX: 0, VY: 0, VZ: 0}

// 	for i := 0; i < 1000; i++ {
// 		io.ApplyGravity(europa, ganymede, callisto)
// 		europa.ApplyGravity(io, ganymede, callisto)
// 		ganymede.ApplyGravity(io, europa, callisto)
// 		callisto.ApplyGravity(io, europa, ganymede)

// 		io.ApplyVelocity(io, europa, ganymede, callisto)

// 		fmt.Println(io, europa, ganymede, callisto)
// 	}

// 	totalEnergy := io.TotalEnergy(io, europa, ganymede, callisto)
// 	fmt.Println("Total Energy:", totalEnergy)
// }

// func PartII() int64 {

// 	io_init := &Moon{X: 3, Y: -6, Z: 6, VX: 0, VY: 0, VZ: 0}
// 	europa_init := &Moon{X: 10, Y: 7, Z: -9, VX: 0, VY: 0, VZ: 0}
// 	ganymede_init := &Moon{X: -3, Y: -7, Z: 9, VX: 0, VY: 0, VZ: 0}
// 	callisto_int := &Moon{X: -8, Y: 0, Z: 4, VX: 0, VY: 0, VZ: 0}

// 	io := &Moon{X: 3, Y: -6, Z: 6, VX: 0, VY: 0, VZ: 0}
// 	europa := &Moon{X: 10, Y: 7, Z: -9, VX: 0, VY: 0, VZ: 0}
// 	ganymede := &Moon{X: -3, Y: -7, Z: 9, VX: 0, VY: 0, VZ: 0}
// 	callisto := &Moon{X: -8, Y: 0, Z: 4, VX: 0, VY: 0, VZ: 0}

// 	var ioEq, euroEq, ganEq, calEq bool

// 	var steps int64

// 	for {
// 		io.ApplyGravity(europa, ganymede, callisto)
// 		europa.ApplyGravity(io, ganymede, callisto)
// 		ganymede.ApplyGravity(io, europa, callisto)
// 		callisto.ApplyGravity(io, europa, ganymede)

// 		io.ApplyVelocity(io, europa, ganymede, callisto)

// 		ioEq = io.IsPositionEqual(io_init) && io.IsVelocityEqual(io_init)
// 		euroEq = europa.IsPositionEqual(europa_init) && europa.IsVelocityEqual(europa_init)
// 		ganEq = ganymede.IsVelocityEqual(ganymede_init)
// 		calEq = callisto.IsPositionEqual(callisto_int) && callisto.IsVelocityEqual(callisto_int)

// 		if ioEq && euroEq && ganEq && calEq {
// 			return steps
// 		}

// 		steps++
// 	}
// }

// func main() {
// 	//PartI()
// 	steps := PartII()
// 	fmt.Printf("God! it took %v steps???. What kept you so long? (╯°□°)╯︵ ┻━┻", steps)
// }
