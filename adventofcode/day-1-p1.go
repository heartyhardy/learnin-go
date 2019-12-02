package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"strconv"
// 	"strings"
// )

// func getstrings() []string {
// 	data, err := ioutil.ReadFile("day-1-input.txt")
// 	if err != nil {
// 		print("Reading data file failed. Cannot skip this criical operation.")
// 		panic("Quitting")
// 	}

// 	moduleMasses := strings.Split(string(data), "\n")

// 	return moduleMasses
// }

// func calModuleMass(modules []string) int {
// 	fuelreq := 0
// 	for i := 0; i < len(modules); i++ {
// 		mass, err := strconv.Atoi(modules[i])
// 		if err != nil {
// 			print("Invalid data format!")
// 			panic("Quitting")
// 		}

// 		fuelreq += int(math.Floor(float64(mass/3))) - 2
// 	}
// 	return fuelreq
// }

// func main() {
// 	modules := getstrings()
// 	fuelreq := calModuleMass(modules)
// 	fmt.Println(fuelreq)
// }
