package main

func isOneBitCharacter(bits []int) bool {
	ones := 0
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 0 {
			break
		}
		ones++
	}
	return ones%2 == 0
}

// func main() {
// 	a := []int{0, 0, 0, 0, 1, 1, 1, 1, 0}
// 	res := isOneBitCharacter(a)
// 	fmt.Println(res)
// }
