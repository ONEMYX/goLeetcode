package main

import "fmt"

// 或与
func solution(cards []int) int {
	// Edit your code here
	result := 0
	for _, v := range cards {
		result ^= v
	}

	return result
}

func main() {
	// Add your test cases here

	fmt.Println(solution([]int{1, 1, 2, 2, 3, 3, 4, 5, 5}) == 4)
	fmt.Println(solution([]int{0, 1, 0, 1, 2}) == 2)
}
