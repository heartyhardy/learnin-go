package main

import "testing"

func TestRotateArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	expected := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, 2)

	for i, v := range nums {
		if v != expected[i] {
			t.Error("Not Equal after rotating")
			break
		}
	}
}
