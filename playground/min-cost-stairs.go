package main

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-2], cost[i-1])
	}
	return min(cost[len(cost)-2], cost[len(cost)-1])
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// func main() {
// 	a := []int{10, 15, 20}
// 	res := minCostClimbingStairs(a)
// 	fmt.Println(res)
// }
