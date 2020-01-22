package aoc15

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type adjacencyMatrix map[string]map[string]int

func readInstructions(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading input file!")
	}
	content := strings.Split(strings.Replace(string(data), "to", "", -1), "\n")
	return content
}

func fillMatrix(content []string) adjacencyMatrix {
	matrix := make(adjacencyMatrix)
	for _, edge := range content {
		fields := strings.FieldsFunc(edge, func(c rune) bool {
			return !unicode.IsNumber(c) && !unicode.IsLetter(c)
		})
		weight, _ := strconv.Atoi(fields[2])
		if _, ok := matrix[fields[0]]; !ok {
			matrix[fields[0]] = make(map[string]int)
		}
		if _, ok := matrix[fields[1]]; !ok {
			matrix[fields[1]] = make(map[string]int)
		}
		matrix[fields[0]][fields[1]] = weight
		matrix[fields[1]][fields[0]] = weight
	}
	return matrix
}

func printGraph(matrix adjacencyMatrix) {
	for k1, v1 := range matrix {
		fmt.Printf("\n * %v --> \n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\n\t-- %-20v  %v", k2, v2)
		}
		fmt.Println()
	}
}

func findMinDistance(matrix adjacencyMatrix) int {
	//start with a random node
	//If its not in the closed list, list all edges
	//Add the current edge to the closed list
	//Pick the min distance from the listed edges
	//Sum up edge weight and compare with the current minDistance
	//repeat until no edges to list
	minDistance := math.MaxInt32
	for vtx := range matrix {
		closed := make(map[string]bool)
		minVtx, distance := vtx, 0

		for !closed[minVtx] {
			minChanged, minEdge := false, math.MaxInt32
			closed[minVtx] = true

			for c, v := range matrix[minVtx] {
				if minEdge > v && !closed[c] {
					minEdge, minVtx, minChanged = v, c, true
				}

			}
			if minChanged {
				distance += minEdge
			}
		}
		if minDistance > distance && distance != 0 {
			minDistance = distance
		}
	}
	return minDistance
}

func findMaxDistance(matrix adjacencyMatrix) int {
	//start with a random node
	//If its not in the closed list, list all edges
	//Add the current edge to the closed list
	//Pick the max distance from the listed edges
	//Sum up edge weight and compare with the current maxDistance
	//repeat until no edges to list
	maxDistance := math.MinInt32
	for vtx := range matrix {
		closed := make(map[string]bool)
		maxVtx, distance := vtx, 0

		for !closed[maxVtx] {
			maxChanged, maxEdge := false, math.MinInt32
			closed[maxVtx] = true

			for c, v := range matrix[maxVtx] {
				if maxEdge < v && !closed[c] {
					maxEdge, maxVtx, maxChanged = v, c, true
				}

			}
			if maxChanged {
				distance += maxEdge
			}
		}
		if maxDistance < distance && distance != 0 {
			maxDistance = distance
		}
	}
	return maxDistance
}

//Run Solution
func Run() {
	content := readInstructions("./inputs/day-09.txt")
	matrix := fillMatrix(content)
	//printGraph(matrix)
	minDistance, maxDistance := findMinDistance(matrix), findMaxDistance(matrix)

	fmt.Println("\n-- AoC 2015: Day 9: All in a Single Night --")
	fmt.Printf("\n🚀  Shortest Path: %20v \n🚂  Longest Path: %21v\n", minDistance, maxDistance)
	fmt.Println("\n-- DONE --")
	fmt.Println("")
}
