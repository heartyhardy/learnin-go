package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readPackageWeights(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func makeCombinations(weightsTable []string, partitionSize int) {
	weights := make([]int, len(weightsTable))
	sum := 0
	for i, w := range weightsTable {
		weight, _ := strconv.Atoi(w)
		weights[i] = weight
		sum += weight
	}
	partition := sum / partitionSize

	for i := 0; i < len(weights); i++ {
		prod := 1
		tmp := make([]int, len(weights))
		copy(tmp, weights)
		tmp = append(tmp[:i], tmp[i+1:]...)
		combinations := populate(tmp, 0, partition)

		for _, v := range combinations {
			prod *= v
		}
		fmt.Println("product", prod)
	}
}

func populate(weights []int, n, size int) (combinations []int) {
	combinations = make([]int, 0)

	for i := len(weights) - 1; i >= 0; i-- {
		if n+weights[i] <= size {
			n += weights[i]
			combinations = append(combinations, weights[i])
		}
	}
	return
}

//Run solution
func Run() {
	weightsTable := readPackageWeights("./inputs/day-24.txt")
	makeCombinations(weightsTable, 3)
}
