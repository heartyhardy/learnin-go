package aoc15

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readFileAndSplitn(filename string) [][]string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Error reading file !")
	}
	rows := strings.Split(string(data), "\n")
	res := make([][]string, len(rows))
	for i, r := range rows {
		res[i] = strings.Split(r, "x")
	}
	return res
}

func calcAll(paper [][]string) (int64, int64) {
	var l, w, h, area, ribbons float64 = 0, 0, 0, 0, 0
	for _, r := range paper {
		l, _ = strconv.ParseFloat(r[0], 64)
		w, _ = strconv.ParseFloat(r[1], 64)
		h, _ = strconv.ParseFloat(r[2], 64)

		lw := 2 * l * w
		wh := 2 * w * h
		hl := 2 * h * l

		//Part I calculate paper
		ma := math.Min(hl, math.Min(lw, wh)) / 2
		area += lw + wh + hl + ma
		mp := math.Min(2*(l+w), math.Min(2*(w+h), 2*(h+l)))
		//Part II calculate ribbons
		bow := l * w * h
		ribbons += mp + bow
	}
	return int64(area), int64(ribbons)
}

// Run solution
// func Run() {
// 	res := readFileAndSplitn("./inputs/day-02.txt")
// 	paper, ribbons := calcAll(res)
// 	fmt.Println("\n-- AoC 2015: Day 02 - I Was Told There Would Be No Math --")
// 	fmt.Printf("\nTotal Paper: %v \nTotal Ribbons: %v\n", paper, ribbons)
// 	fmt.Println("\n-- DONE --")
// 	fmt.Println("")
// }
