package aoc15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type ingrediantsTable []*ingrediant

type ingrediant struct {
	name       string
	capacity   int
	texture    int
	durability int
	flavor     int
	calories   int
}

func readIngrediatns(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func makeIngrediantsTable(ingrediants []string) ingrediantsTable {
	table := make([]*ingrediant, len(ingrediants))
	for i, ing := range ingrediants {
		props := strings.FieldsFunc(ing, func(c rune) bool {
			return !unicode.IsNumber(c) && c != rune('-')
		})
		newIng := new(ingrediant)
		newIng.name = strings.Split(ing, ":")[0]
		newIng.capacity, _ = strconv.Atoi(props[0])
		newIng.durability, _ = strconv.Atoi(props[1])
		newIng.flavor, _ = strconv.Atoi(props[2])
		newIng.texture, _ = strconv.Atoi(props[3])
		newIng.calories, _ = strconv.Atoi(props[4])

		table[i] = newIng
	}
	return table
}

func makeRecipe(table ingrediantsTable, amounts []int) (int, int) {
	var capacity, durability, flavor, texture, calories int
	for k, v := range table {
		capacity += v.capacity * amounts[k]
		durability += v.durability * amounts[k]
		flavor += v.flavor * amounts[k]
		texture += v.texture * amounts[k]
		calories += v.calories * amounts[k]
	}
	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}
	return calories, capacity * durability * flavor * texture
}

func shakeNMake(table ingrediantsTable, isScoreBased bool) int {
	// i = sprinkles j = peanut butter k = frosting l = sugar
	amounts, maxScore := make([]int, 5), 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			for k := 1; k <= 100; k++ {
				for l := 1; l <= 100; l++ {
					if i+j+k+l == 100 {
						amounts[0], amounts[1], amounts[2], amounts[3] = i, j, k, l
						if isScoreBased {
							if _, score := makeRecipe(table, amounts); maxScore < score {
								maxScore = score
							}
							continue
						}
						if calories, score := makeRecipe(table, amounts); calories == 500 && maxScore < score {
							maxScore = score
						}
					}
				}
			}
		}
	}
	return maxScore
}

//Run solution
func Run() {
	ingrediants := readIngrediatns("./inputs/day-15.txt")
	table := makeIngrediantsTable(ingrediants)
	scoreBased := shakeNMake(table, true)
	caloriesBased := shakeNMake(table, false)

	fmt.Printf("\n\t -- Day 15: Science for Hungry People --\n")
	fmt.Printf("\n🍔\tMax Score Recipe: %25v \n🥗\tBest Calories Recipe: %21v\n", scoreBased, caloriesBased)
	fmt.Println(`
	______
	|  _  \
	| | | |___  _ __   ___
	| | | / _ \| '_ \ / _ \
	| |/ / (_) | | | |  __/
	|___/ \___/|_| |_|\___|`)
	fmt.Println("")
}
