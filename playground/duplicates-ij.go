package main

func containsNearbyDuplicate(nums []int, k int) bool {
	hm := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if val, ok := hm[nums[i]]; ok && i-val <= k {
			return true
		} else {
			hm[nums[i]] = i
		}
	}
	return false
}

// func main() {
// 	a := []int{1, 0, 1, 1}
// 	res := containsNearbyDuplicate(a, 1)
// 	fmt.Println(res)
// }
