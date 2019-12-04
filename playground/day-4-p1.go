package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func checkRules(input int) bool {
// 	reps := map[int]int{}
// 	fields := []string{}
// 	i := strconv.Itoa(input)
// 	fields = strings.Split(i, "")

// 	isLower, isSame := false, false
// 	var s1, s2 int = 0, 0

// 	for j := 0; j < len(fields)-1; j++ {
// 		s1, _ = strconv.Atoi(fields[j])
// 		s2, _ = strconv.Atoi(fields[j+1])

// 		if s2 > s1 {
// 			isLower = false
// 		} else if s2 < s1 {
// 			isLower = true
// 			break
// 		} else {
// 			reps[s1] = reps[s1] + 1
// 			// if isSame != true {
// 			// 	isSame = true
// 			// }
// 		}
// 	}

// 	for _, v := range reps {
// 		if v == 1 {
// 			isSame = true
// 		}
// 	}

// 	return isSame && !isLower
// }

// func genPass() int {
// 	lower := 125730
// 	upper := 579381

// 	cnt := 0

// 	for i := lower; i <= upper; i++ {
// 		isvalid := checkRules(i)
// 		if isvalid {
// 			cnt++
// 		}
// 	}

// 	return cnt
// }

// func main() {
// 	res := genPass()
// 	fmt.Println(res)

// }
