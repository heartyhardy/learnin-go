package sorting

//InsertionSort ...
func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		k := nums[i]
		j := i - 1

		for j >= 0 && k < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = k
	}
	return nums
}

//MergeSort ...
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	m := len(nums) / 2
	return Merge(MergeSort(nums[:m]), MergeSort(nums[m:]))
}

//Merge ...
func Merge(r, l []int) []int {
	size := len(r) + len(l)
	res := make([]int, size, size)

	var i, j int
	for k := 0; k < size; k++ {
		switch {
		case i == len(l):
			res[k] = r[j]
			j++
		case j == len(r):
			res[k] = l[i]
			i++
		case l[i] > r[j]:
			res[k] = r[j]
			j++
		case l[i] < r[j]:
			res[k] = l[i]
			i++
		case l[i] == r[j]:
			res[k] = r[j]
			j++
		}
	}
	return res
}
