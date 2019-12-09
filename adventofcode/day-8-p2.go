package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func readRawData(file string) map[int][]int {
// 	layerSize := 25 * 6
// 	data, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	strPixels := strings.Split(strings.TrimSpace(string(data)), "")

// 	layers := map[int][]int{}
// 	pixels := make([]int, len(strPixels))
// 	layerId := 0

// 	for i := 0; i < len(strPixels); i++ {
// 		pixels[i], _ = strconv.Atoi(strPixels[i])
// 	}

// 	for i := 0; i < len(pixels); i = i + layerSize {
// 		layers[layerId] = pixels[i : i+150]
// 		layerId++
// 	}
// 	return layers
// }

// func countNumber(layer []int, num int) int {
// 	nums := 0
// 	for _, pixel := range layer {
// 		if pixel == num {
// 			nums++
// 		}
// 	}
// 	return nums
// }

// func seekMin(layers map[int][]int) (minID int, minCount int) {
// 	minCount, minID = 999, -1
// 	for id, layer := range layers {
// 		current := countNumber(layer, 0)
// 		if minCount > current {
// 			minCount = current
// 			minID = id
// 		}
// 	}
// 	return
// }

// func partI(layerId int, layers map[int][]int) int {
// 	ones := countNumber(layers[layerId], 1)
// 	twos := countNumber(layers[layerId], 2)
// 	return ones * twos
// }

// func decode(mergedLayer []int, width int) {
// 	fmt.Println()
// 	for i := 0; i < len(mergedLayer); i++ {
// 		if mergedLayer[i] != 0 {
// 			fmt.Printf("\u001b[33m\u001b[44m%v", "☰") //Replaced 1 here with the fancy symbol
// 		} else {
// 			fmt.Print("\u001b[40m⋰")
// 		}
// 		if (i+1)%width == 0 {
// 			fmt.Print("\u001b[0m\n")
// 		}
// 	}
// 	fmt.Println()
// }

// func mergeLayers(layers map[int][]int) []int {
// 	rowResult := layers[0]
// 	for id := 1; id < len(layers); id++ {
// 		for pixelId, pixel := range layers[id] {
// 			switch rowResult[pixelId] {
// 			case 2:
// 				rowResult[pixelId] = pixel
// 			}
// 		}
// 	}
// 	return rowResult
// }

// func main() {
// 	layers := readRawData("day-8-input.txt")

// 	merged := mergeLayers(layers)
// 	decode(merged, 25)
// }
