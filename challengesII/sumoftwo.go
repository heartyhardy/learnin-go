package main

import "fmt"

func sumoftwo(arr []int, sum int) (i1 int, i2 int, ok bool) {
	m := map[int]int{}

	for i := 0; i < len(arr); i++ {
		if v, _ok := m[arr[i]]; _ok == true {
			i1 = v
			i2 = i
			ok = true
			return
		}
		m[sum-arr[i]] = i
	}
	fmt.Println(m)
	return
}

func main() {
	a := []int{1, 8, 5, 7, 10, 9}
	i1, i2, _ok := sumoftwo(a, 14)
	fmt.Printf("[%v, %v] result: %v\n", i1, i2, _ok)
}
