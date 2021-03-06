package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"strings"
// )

// const (
// 	space   = "."
// 	astroid = "#"
// )

// type coords struct {
// 	x    int
// 	y    int
// 	los  int
// 	dist float64
// }

// func parse(input string) []coords {
// 	strs := strings.Split(strings.TrimSpace(input), "\n")
// 	mapdata := []coords{}
// 	for r, row := range strs {
// 		for c, col := range row {
// 			if string(col) == astroid {
// 				newcoords := coords{
// 					x:   c,
// 					y:   r,
// 					los: 0,
// 				}
// 				mapdata = append(mapdata, newcoords)
// 			}
// 		}
// 	}
// 	return mapdata
// }

// func angle(x, y int) float64 {
// 	return math.Atan2(float64(y), float64(x)) * 180 / math.Pi
// }

// func clockwiseAngle(atan2 float64) float64 {
// 	switch {
// 	case atan2 >= 90.0 && atan2 <= 180:
// 		return atan2 - 90.0
// 	case atan2 < -180 && atan2 <= -90:
// 		return 270 + atan2
// 	case atan2 > -90 && atan2 <= 0:
// 		return 270 + atan2
// 	case atan2 >= 0 && atan2 < 90:
// 		return 360 - atan2
// 	}
// 	return atan2
// }

// func lineofsight(mapdata []coords) []coords {
// 	var currcoord, xdiff, ydiff int
// 	losdata := map[float64]coords{}
// 	for currcoord < len(mapdata) {
// 		for _, coord := range mapdata {
// 			xdiff = mapdata[currcoord].x - coord.x
// 			ydiff = mapdata[currcoord].y - coord.y
// 			atan2 := angle(xdiff, ydiff)
// 			if _, ok := losdata[atan2]; !ok {
// 				losdata[atan2] = coord
// 				mapdata[currcoord].los++
// 			}
// 		}
// 		currcoord++

// 		for k := range losdata {
// 			delete(losdata, k)
// 		}
// 	}
// 	return mapdata
// }

// func buildAstroidCache(astroids []coords, astroidID int) (map[float64]coords, map[float64]float64) {
// 	var xdiff, ydiff int
// 	astroidCache := map[float64]coords{}
// 	astroidDist := map[float64]float64{}
// 	for _, astroid := range astroids {
// 		xdiff = astroids[astroidID].x - astroid.x
// 		ydiff = astroid.y - astroids[astroidID].y
// 		dist := math.Sqrt(float64(xdiff)*float64(xdiff) + float64(ydiff)*float64(ydiff))
// 		astroid.dist = dist
// 		atan2 := angle(xdiff, ydiff)
// 		if atan2 < 0 {
// 			atan2 = atan2 + 2*math.Pi
// 		}
// 		if _, ok := astroidCache[atan2]; !ok {
// 			astroidCache[atan2] = astroid
// 		}
// 		astroidDist[dist] = atan2
// 	}
// 	return astroidCache, astroidDist
// }

// func partI(data []byte) ([]coords, int) {
// 	maxlos, astroidID := 0, 0
// 	bestlookout := coords{}
// 	coords := parse(string(data))
// 	astroids := lineofsight(coords)

// 	for i, astroid := range astroids {
// 		if maxlos < astroid.los {
// 			maxlos = astroid.los
// 			bestlookout = astroid
// 			astroidID = i
// 		}
// 	}

// 	fmt.Println("\n\u001b[34m------------------- BEST LINE OF SIGHT ---------------------\u001b[0m")
// 	fmt.Printf("\t\t \u001b[32mLOS-INDEX: %v\u001b[0m                   \n", maxlos)
// 	fmt.Printf("\t\t \u001b[32mCOORDINATES: ( %v,%v )\u001b[0m                   \n", bestlookout.x, bestlookout.y)
// 	fmt.Println("\u001b[34m------------------------------------------------------------\u001b[0m")
// 	fmt.Println()

// 	return astroids, astroidID
// }

// func partII(astroids []coords, astroidID int) {
// 	astroidCache, astroDist := buildAstroidCache(astroids, astroidID)

// }

// func main() {
// 	data, _ := ioutil.ReadFile("day-10-input.txt")
// 	astroids, astroidID := partI(data)

// 	partII(astroids, astroidID)
// }
