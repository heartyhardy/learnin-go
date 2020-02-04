package aoc15

import (
	"fmt"
	"math"
	"sort"
	"time"
)

const maxGifts = 33100000

func getDivisors(n int) (divisors []int) {
	i := 1
	divisors = make([]int, 0)
	for i <= int(math.Sqrt(float64(n))) {
		if n%i == 0 {
			if n/i == i {
				tmp := make([]int, len(divisors)+1)
				copy(tmp, divisors)
				tmp[len(tmp)-1] = i
				divisors = make([]int, len(tmp))
				copy(divisors, tmp)
			} else {
				tmp := make([]int, len(divisors)+2)
				copy(tmp, divisors)
				tmp[len(tmp)-1] = i
				tmp[len(tmp)-2] = n / i
				divisors = make([]int, len(tmp))
				copy(divisors, tmp)
			}
		}
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(divisors)))
	return
}

func getGiftCount(divisors *[]int, gph int) int {
	count := 0
	for _, v := range *divisors {
		count += v * gph
	}
	return count
}

func getGiftCountLimited(divisors *[]int, gph, n, maxn int) int {
	count := 0
	for _, v := range *divisors {
		if v*maxn < n {
			break
		}
		count += v * gph
	}
	return count
}

func giftsAway(gph int) {
	i := maxGifts / 1000
	for {
		divisors := getDivisors(i)
		count := getGiftCount(&divisors, gph)

		if count >= maxGifts {
			fmt.Println("Part I: House No: ", i)
			break
		}
		i++
	}
}

func giftsAwayLimited(gph int) {
	i := maxGifts / 1000
	for {
		divisors := getDivisors(i)
		count := getGiftCountLimited(&divisors, gph, i, 50)

		if count >= maxGifts {
			fmt.Println("Part II - Limited: House No: ", i)
			break
		}
		i++
	}
}

func showBusyCursor() {
	for {
		cursors := `\|/-`
		for _, c := range cursors {
			fmt.Printf("\r %v", string(c))
			time.Sleep(100 * time.Millisecond)
		}
	}
}

//Run Solution
func Run() {
	go showBusyCursor()
	go giftsAway(10)
	go giftsAwayLimited(11)
	fmt.Scanln()
}
