package main

import (
	"fmt"
)

func solution(n int, k int, data []int) int {
	// Edit your code here
	currentStock := 0
	num := 0
	for i := 0; i < n; i++ {
		if currentStock > 0 {
			currentStock--
		}
		minData := data[i]
		minIndex := i
		for j := i + 1; j < min(i+k, n); j++ {
			if data[j] < minData {
				minIndex = j
				minData = data[j]
			}
		}
		if minIndex == i {
			buy := min(k-currentStock, n-i-currentStock)
			num += buy * data[minIndex]
			currentStock += buy
		} else {
			buy := max(0, minIndex-i-currentStock)
			num += buy * data[minIndex]
			currentStock += buy
		}
	}
	return num
}

func main() {
	// Add your test cases here

	fmt.Println(solution(5, 2, []int{1, 2, 3, 3, 2}) == 9)
}
