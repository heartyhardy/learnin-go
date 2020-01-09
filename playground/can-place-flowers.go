package main

func canPlaceFlowersII(flowerbed []int, n int) bool {
	sum := 0
	for i, v := range flowerbed {
		if i == 0 && v == 0 {
			flowerbed[i] = 2
			sum++
			continue
		} else if i == 0 && v == 1 {
			continue
		}
		if v == 1 && flowerbed[i-1] == 2 {
			sum--
		} else if v == 0 && flowerbed[i-1] == 0 {
			sum++
			flowerbed[i] = 2
		}
	}
	return sum >= n
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	//sum := len(flowerbed)/2 + len(flowerbed)%2
	vc := 0
	for i := 0; i < len(flowerbed); i++ {
		v := flowerbed[i]
		if v == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				vc++
			}
		}
	}
	return vc >= n
}

// func main() {
// 	a := []int{1, 0, 1}
// 	res := canPlaceFlowers(a, 0)
// 	fmt.Println(res)
// }
