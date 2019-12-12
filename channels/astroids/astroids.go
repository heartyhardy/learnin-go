package astroids

import (
	"io/ioutil"
	"strings"
)

//Astroid ...
const Astroid rune = '#'

//Coordinate ...
type Coordinate struct {
	X int
	Y int
}

//ReadFile ...
func ReadFile(filename string) []Coordinate {
	data, _ := ioutil.ReadFile(filename)
	sinput := strings.Split(string(data), "\n")
	var astroids []Coordinate

	for y, line := range sinput {
		line = strings.TrimSpace(line)

		for x, r := range line {
			if r == Astroid {
				astroids = append(astroids, Coordinate{X: x, Y: y})
			}
		}
	}
	return astroids
}
