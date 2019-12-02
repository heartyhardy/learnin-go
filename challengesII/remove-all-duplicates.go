package main

// import "fmt"

// func removeDuplicates(el []int) int {
// 	i, r := 0, len(el)
// 	for i < len(el)-1 { //O(n ^ 2)
// 		if el[i] > el[i+1] {
// 			t := el[i]
// 			el[i] = el[i+1]
// 			el[i+1] = t
// 			i = 0
// 		} else {
// 			i++
// 		}
// 	}

// 	for i = 0; i < len(el)-1; i++ { // O(n)
// 		if el[i]^el[i+1] == 0 {
// 			if len(el)-1-(i+3) >= 0 {
// 				t1 := el[i+2]
// 				t2 := el[i+3]
// 				el[i+3] = el[i+1]
// 				el[i+2] = el[i]
// 				el[i] = t1
// 				el[i+1] = t2
// 				r -= 2
// 			} else {
// 				r -= 2
// 				break
// 			}
// 		}
// 	}
// 	return r
// }

// func main() {
// 	b := 1 << 1000
// 	fmt.Println(b)
// 	a := []int{2, 5, 4, 4, 5, 3, 2, 3, 1}
// 	res := removeDuplicates(a)
// 	fmt.Println(res)
// }
