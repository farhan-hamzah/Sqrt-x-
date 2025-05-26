package main

import (
	"fmt"
)

// Fungsi untuk menghitung akar kuadrat (integer saja)
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	var ans int

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func main() {
	// Contoh penggunaan
	testCases := []int{0, 1, 4, 8, 10, 15, 16, 100, 2147395599}
	for _, x := range testCases {
		fmt.Printf("mySqrt(%d) = %d\n", x, mySqrt(x))
	}
}
