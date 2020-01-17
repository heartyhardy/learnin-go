package main

import (
	"fmt"
	"math/rand"
)

func bubblesort20200116(nums []int) []int {
	swapped := false
	i, imax, swpn := 0, len(nums)-1, 0
	for i < imax {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
			swapped = true
			swpn = i
		}
		if i == imax-1 && !swapped {
			break
		} else if i == imax-1 && swapped {
			i = -1
			swapped = false
			imax = swpn
		}
		i++
	}
	return nums
}

func selectionSort20200117(nums []int) []int {
	min := 0
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

func shuffle(nums []int) {
	src := rand.NewSource(100)
	rnd := rand.New(src)
	for i := 0; i < len(nums); i++ {
		idx := rnd.Intn(len(nums) - 1)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
	fmt.Println("\nShuffled: ", nums)
}

func mergesort20200117(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	m := len(nums) / 2
	return merge20200117(mergesort20200117(nums[:m]), mergesort20200117(nums[m:]))
}

func merge20200117(l, r []int) []int {
	size, i, j := len(l)+len(r), 0, 0
	res := make([]int, len(l)+len(r))

	for k := 0; k < size; k++ {
		switch {
		case i == len(l):
			res[k] = r[j]
			j++
		case j == len(r):
			res[k] = l[i]
			i++
		case l[i] < r[j]:
			res[k] = l[i]
			i++
		case l[i] > r[j]:
			res[k] = r[j]
			j++
		case l[i] == r[j]:
			res[k] = r[j]
			j++
		}
	}
	return res
}

func partition20200117(nums []int) int {
	pivot := nums[len(nums)-1]
	pindex := len(nums) - 1
	l, r := 0, pindex
	for l <= r {
		if nums[l] >= pivot {
			nums[pindex], nums[pindex-1] = nums[pindex-1], nums[pindex]
			if pindex-l > 1 {
				nums[pindex], nums[l] = nums[l], nums[pindex]
			}
			pindex--
			r--
		}
		l++
	}
	return pindex
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	pivot := rand.Int() % len(nums)
	//fmt.Println(pivot, left, right, nums[left], nums[right], nums[pivot])

	nums[pivot], nums[right] = nums[right], nums[pivot]

	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	quicksort(nums[:left])
	quicksort(nums[left+1:])

	return nums
}

func main() {
	nums := []int{1, 5, 3, 2, 7, 5, 7, 9, 6, 11, 10}
	res := bubblesort20200116(nums)
	fmt.Println("\n -- Bubble sort --")
	fmt.Println(res)
	shuffle(nums)
	res = selectionSort20200117(nums)
	fmt.Println("\n -- Selection sort --")
	fmt.Println(res)
	shuffle(nums)
	res = mergesort20200117(nums)
	fmt.Println("\n -- Merge sort --")
	fmt.Println(res)
	shuffle(nums)
	fmt.Println("\nQuick sort")
	res = quicksort(nums)
	fmt.Println(res)
}
