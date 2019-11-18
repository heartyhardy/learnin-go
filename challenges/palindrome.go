// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func isPalindrome(word string) bool {
// 	arr := []byte(strings.ToLower(word))
// 	res := false

// 	for i := 0; i < len(arr)/2; i++ {
// 		if arr[i] != arr[len(arr)-1-i] {
// 			res = false
// 			break
// 		} else {
// 			res = true
// 		}
// 	}
// 	return res
// }

// func main() {
// 	res := isPalindrome("Madam")
// 	fmt.Printf("Is palindrome: %v\n", res)
// }
