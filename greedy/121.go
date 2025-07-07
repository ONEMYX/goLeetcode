package main

import "fmt"

func maxProfits(prices []int) int {
	mx := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			mx = max(mx, prices[i]-prices[j])
		}
	}
	return mx
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := maxProfits(list)
	fmt.Println(s)
}
