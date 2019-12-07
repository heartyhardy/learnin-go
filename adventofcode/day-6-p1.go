package main

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func readMapData(file string) []string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		os.Exit(1)
	}
	dataNodes := strings.Split(string(data), "\n")
	sort.Strings(dataNodes)

	return dataNodes
}

func saveOrbits(dataNodes []string) map[string]string {
	orbitData := map[string]string{}
	for _, node := range dataNodes {
		planetData := strings.Split(node, ")")
		parent := planetData[0]
		child := planetData[1]
		_, hasParent := orbitData[parent]
		_, hasChild := orbitData[child]

		if !hasParent && !hasChild {
			orbitData[child] = parent
			orbitData[parent] = "nil"
		} else if !hasParent && hasChild {
			orbitData[child] = parent
			orbitData[parent] = "nil"
		} else if hasParent && !hasChild {
			orbitData[child] = parent
		} else if hasParent && hasChild {
			orbitData[child] = parent
		}
	}

	return orbitData
}

func calculateOrbits(orbitData map[string]string) int {
	orbits := 0
	current := ""

	for _, pointer := range orbitData {
		current = pointer
		for current != "" && current != "nil" {
			orbits += 1
			current = orbitData[current]
		}
	}
	return orbits
}

func calculateOrbitalTransfers(orbitData map[string]string) int {
	santaPath := map[string]string{}
	yourPath := map[string]string{}
	transfers := 0

	start := orbitData["YOU"]

	current := start

	for current != "" && current != "nil" {
		santaPath[current] = orbitData[current]
		current = orbitData[current]
	}

	start = orbitData["SAN"]

	current = start

	for current != "" && current != "nil" {
		yourPath[current] = orbitData[current]
		current = orbitData[current]
	}

	for _, path := range santaPath {
		if _, hasPath := yourPath[path]; !hasPath {
			transfers++
		}
	}

	for _, path := range yourPath {
		if _, hasPath := santaPath[path]; !hasPath {
			transfers++
		}
	}

	return transfers
}

// func main() {
// 	res := readMapData("day-6-input.txt")
// 	orbitNodes := saveOrbits(res)
// 	orbits := calculateOrbits(orbitNodes)
// 	fmt.Println("Number of Orbits: ", orbits)

// 	transfers := calculateOrbitalTransfers(orbitNodes)
// 	fmt.Println("Minimum number of Transfers: ", transfers)
// }
