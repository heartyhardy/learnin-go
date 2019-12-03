package main

import "fmt"

func plusOne(digits []int) []int {
	// important one in the head
	// has to add to the tail
	// if last digit is 9 adding 1 will carry 1 to the last - 1 array pos
	// 999 + 1 = ?
	target := digits[len(digits)-1]
	if target != 9 {
		digits[len(digits)-1] = target + 1
		return digits
	}

	i := len(digits) - 1
	carry := false

	for i >= 0 {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			carry = false
			break
		}
		digits[i] = 0
		carry = true
		i--
		continue
	}

	if carry == true {
		digits = append(digits, 0)
		digits[0] = 1
	}

	return digits
}

func main() {
	a := []int{9, 9}
	res := plusOne(a)
	fmt.Println(res)
}
