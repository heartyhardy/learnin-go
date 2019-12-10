package main

import (
	"math"
	"sort"
	"strings"
)

//Coordinate ...
type Coordinate struct {
	X, Y int
}

// AbsInt ...
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MaxInt ...
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//Distance ...
func (c Coordinate) Distance(t Coordinate) int {
	return AbsInt(c.X-t.X) + AbsInt(c.Y-t.Y)
}

//Angle ...
func (c Coordinate) Angle(t Coordinate) float64 {
	angle := math.Atan2(float64(t.X-c.X), float64(c.Y-t.Y)) * 180 / math.Pi
	if angle < 0 {
		angle = angle + 360
	}
	return angle
}

const (
	asteroid = '#'
)

type void struct{}

//AsteroidMap ...
func AsteroidMap(input []string) []Coordinate {
	var belt []Coordinate
	for y, line := range input {
		line = strings.TrimSpace(line)

		for x, r := range line {
			if r == asteroid {
				belt = append(belt, Coordinate{X: x, Y: y})
			}
		}
	}
	return belt
}

//BestLocation ...
func BestLocation(input []string) (Coordinate, int) {
	belt := AsteroidMap(input)
	var bestCount int
	var best Coordinate

	for _, current := range belt {
		angles := make(map[float64]void)
		var count int

		for _, target := range belt {
			if current == target {
				continue
			}

			angle := current.Angle(target)

			if _, ok := angles[angle]; !ok {
				angles[angle] = void{}
				count++
			}
		}

		if count > bestCount {
			bestCount = count
			best = current
		}
	}

	return best, bestCount
}

//CompleteVaporizationBet ...
func CompleteVaporizationBet(input []string, station Coordinate, bet int) int {
	belt := AsteroidMap(input)

	targets := make(map[float64]map[int]Coordinate)
	var angles []float64

	for _, target := range belt {
		if target == station {
			continue
		}

		angle := station.Angle(target)
		distance := station.Distance(target)

		if _, ok := targets[angle]; !ok {
			targets[angle] = make(map[int]Coordinate)
			angles = append(angles, angle)
		}
		targets[angle][distance] = target
	}

	sort.Float64s(angles)

	var hit int

	for len(angles) > 0 {
		for _, angle := range angles {
			closest := -1

			for distance := range targets[angle] {
				if distance < closest || closest == -1 {
					closest = distance
				}
			}

			hit++
			if hit == bet {
				return targets[angle][closest].X*100 + targets[angle][closest].Y
			}

			delete(targets[angle], closest)
		}

		var clean []float64
		for _, angle := range angles {
			if len(targets[angle]) > 0 {
				clean = append(clean, angle)
			}
		}
		angles = clean
	}

	return -1
}
