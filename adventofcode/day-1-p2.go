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

// // While loop solution
// func calExtraFuel(moduleFuel int) int {
// 	exfuel := 0
// 	m := moduleFuel
// 	for m > 0 {
// 		m = int(math.Floor(float64(m/3))) - 2
// 		if m > 0 {
// 			exfuel += m
// 		}
// 	}
// 	return exfuel
// }

// // Recursive solution
// func calExtraFuelRec(moduleFuel int) int {
// 	exfuel := 0
// 	m := int(math.Floor(float64(moduleFuel/3))) - 2

// 	if m <= 0 {
// 		return 0
// 	}
// 	return m + calExtraFuel(m)

// 	return exfuel
// }

// // Calculate All Fuel requirements
// func calModuleMass(modules []string) int {
// 	fuelreq := 0
// 	for i := 0; i < len(modules); i++ {
// 		mass, err := strconv.Atoi(modules[i])
// 		if err != nil {
// 			print("Invalid data format!")
// 			panic("Quitting")
// 		}

// 		fuelreq += calExtraFuelRec(mass)
// 	}
// 	return fuelreq
// }

// func main() {
// 	modules := getstrings()
// 	fuelreq := calModuleMass(modules)
// 	fmt.Println(fuelreq)
// }
