package aoc15

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func genMd5(key string, zeros int) int {
	i := 0
	for {
		newSeed := []byte(fmt.Sprintf("%s%s", key, strconv.Itoa(i)))
		md5seed := md5.Sum(newSeed)
		prefix := fmt.Sprintf("%x", md5seed)[:zeros]
		v, err := strconv.Atoi(prefix)
		if err == nil && v == 0 {
			return i
		}
		i++
	}
}

//Run Solution
// func Run() {
// 	fiveZeros := genMd5("ckczppom", 5)
// 	sixZeros := genMd5("ckczppom", 6)
// 	fmt.Println("\n-- AoC 2015: Day 4: The Ideal Stocking Stuffer --")
// 	fmt.Printf("\nFirst occurrence of 5 leading Zeros: %v \nFirst Basement Occurrence: %v", fiveZeros, sixZeros)
// 	fmt.Println("\n-- DONE --")
// 	fmt.Println("")
// }
