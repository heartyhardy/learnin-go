package main

import (
	"fmt"
	"math"
)

func findMaxAverageIII(nums []int, k int) float64 {
	curavg, maxavg := 0., float64(math.MinInt32)
	for i := 0; i < len(nums)-k+1; i++ {
		for j := i; j < i+k; j++ {
			curavg += float64(nums[j])
		}
		if maxavg < curavg/float64(k) {
			maxavg = curavg / float64(k)
		}
		curavg = 0.
	}
	return maxavg
}

func findMaxAverage(nums []int, k int) float64 {
	sum, max := 0, 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max = sum

	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

func main() {
	a := []int{1, 12, -5, -6, 50, 3}
	k := 4
	res := findMaxAverage(a, k)
	fmt.Println(res)
}
